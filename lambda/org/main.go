package main

import (
	"context"
	"encoding/json"
	"errors"
	"nippo/cognito"
	"nippo/database"
	"nippo/database/tidb"
	"nippo/input"
	"nippo/lambda/response"
	"nippo/log"
	"nippo/schema"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"gorm.io/gorm"
)

type variable struct {
	userPool string
}

var env variable
var db database.Database
var cognitoClient *cognitoidentityprovider.Client

func cacheEnv() error {
	if env.userPool = os.Getenv("USER_POOL"); len(env.userPool) == 0 {
		return errors.New("env for user pool is not set")
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

	cognitoClient = cognitoidentityprovider.NewFromConfig(cfg)

	return nil
}

func handler(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	claims, ok := req.RequestContext.Authorizer["claims"].(map[string]any)
	if !ok {
		return response.BadRequest(), nil
	}

	raw, ok := claims["email"].(string)
	if !ok {
		return response.BadRequest(), nil
	}

	email, err := input.ValidEmail(raw)
	if err != nil {
		log.Warn(err)
		return response.BadRequest(), nil
	}

	orgId, err := db.ReadUserOrgId_Email(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Warnf("%s has token but not found on database", email)
		return response.BadRequest(), nil
	} else if err != nil {
		log.Error(err)
		return response.Internal(), nil
	}

	err = cognito.MarkEmailChecked(ctx, cognitoClient, env.userPool, email)
	if err != nil {
		log.Error(err)
		return response.Internal(), nil
	}

	authOrg := schema.AuthOrg{
		OrgId: orgId,
	}
	b, err := json.Marshal(authOrg)
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

	if err := configAws(); err != nil {
		log.Fatal(err)
	}

	if err := connectDb(); err != nil {
		log.Fatal(err)
	}

	lambda.Start(handler)
}
