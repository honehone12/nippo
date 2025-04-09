package main

import (
	"context"
	"errors"
	"nippo/database"
	"nippo/database/tidb"
	"nippo/lambda/response"
	"nippo/log"
	"nippo/view/q"
	"nippo/view/viewprops"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type variable struct {
	orgIdStr string
	dbUri    string
	bucket   string
}

var (
	env      variable
	db       database.Database
	s3Client *s3.Client
)

func collectEnv() error {
	if env.orgIdStr = os.Getenv("ORG_ID"); len(env.orgIdStr) == 0 {
		return errors.New("env for orgIdStr is not set")
	}

	if env.dbUri = os.Getenv("DB_URI"); len(env.dbUri) == 0 {
		return errors.New("env for dbUri is not set")
	}

	if env.bucket = os.Getenv("BUCKET"); len(env.bucket) == 0 {
		return errors.New("env for bucket is not set")
	}

	return nil
}

func connectDb() error {
	var err error
	db, err = tidb.Open(env.dbUri)
	return err
}

func configAws() error {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return err
	}

	s3Client = s3.NewFromConfig(cfg)
	return nil
}

func handler(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	params := req.QueryStringParameters

	q, err := q.FromParams(params)
	if err != nil {
		return response.BadRequest(), nil
	}

	props := viewprops.NewProps(
		ctx,
		env.bucket,
		env.orgIdStr,
		db,
		s3Client,
	)

	b, err := q.Query(props)
	if err != nil {
		log.L.Error(err)
		return response.Internal(), nil
	}

	return response.Json(b), nil
}

func main() {
	log.FmtPlain()

	if err := collectEnv(); err != nil {
		log.L.Fatal(err)
	}

	if err := connectDb(); err != nil {
		log.L.Fatal(err)
	}

	if err := configAws(); err != nil {
		log.L.Fatal(err)
	}

	lambda.Start(handler)
}
