package main

import (
	"bytes"
	"context"
	"errors"
	"nippo/database"
	"nippo/database/tidb"
	"nippo/input"
	"nippo/lambda/response"
	"nippo/log"
	"nippo/view/q"
	"nippo/view/viewprops"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type variable struct {
	orgId  uuid.UUID
	bucket string
}

var (
	env       variable
	db        database.Database
	s3Client  *s3.Client
	s3Presign *s3.PresignClient
)

func cacheEnv() error {
	orgIdStr := os.Getenv("ORG_ID")
	if len(orgIdStr) == 0 {
		return errors.New("env for orgIdStr is not set")
	}
	orgId, err := uuid.Parse(orgIdStr)
	if err != nil {
		return err
	}
	env.orgId = orgId

	if env.bucket = os.Getenv("BUCKET"); len(env.bucket) == 0 {
		return errors.New("env for bucket is not set")
	}

	return nil
}

func connectDb() error {
	dbUri := os.Getenv("DB_URI")
	if len(dbUri) == 0 {
		return errors.New("env for dbUri is not set")
	}

	var err error
	db, err = tidb.Open(dbUri)
	return err
}

func configAws() error {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return err
	}

	s3Client = s3.NewFromConfig(cfg)
	s3Presign = s3.NewPresignClient(s3Client)
	return nil
}

func getViewerOrgId(auth map[string]any) ([]uint8, error) {
	claims, ok := auth["claims"].(map[string]any)
	if !ok {
		return nil, errors.New("auth not found")
	}

	raw, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("user email not found")
	}

	email, err := input.ValidEmail(raw)
	if err != nil {
		return nil, err
	}

	orgId, err := db.ReadUserOrgIdEm(email)
	if err != nil {
		return nil, err
	}

	return orgId, nil
}

func handler(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	viewerOrgId, err := getViewerOrgId(req.RequestContext.Authorizer)
	if err != nil {
		log.Warn(err)
		return response.BadRequest(), nil
	}

	if !bytes.Equal(viewerOrgId, env.orgId[:]) {
		log.Warn("viewing user has wrong org id")
		return response.BadRequest(), nil
	}

	params := req.QueryStringParameters

	q, err := q.FromParams(params)
	if err != nil {
		return response.BadRequest(), nil
	}

	props := viewprops.NewProps(
		ctx,
		env.bucket,
		env.orgId,
		db,
		s3Client,
		s3Presign,
	)

	b, err := q.Query(props)
	if err != nil {
		log.Error(err)
		return response.Internal(), nil
	}

	return response.Json(b), nil
}

func main() {
	log.FmtPlain()

	if err := cacheEnv(); err != nil {
		log.Fatal(err)
	}

	if err := connectDb(); err != nil {
		log.Fatal(err)
	}

	if err := configAws(); err != nil {
		log.Fatal(err)
	}

	lambda.Start(handler)
}
