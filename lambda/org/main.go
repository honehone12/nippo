package main

import (
	"context"
	"encoding/base64"
	"errors"
	"nippo/database"
	"nippo/database/tidb"
	"nippo/input"
	"nippo/lambda/response"
	"nippo/log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var db database.Database

func connectDb() error {
	dbUri := os.Getenv("DB_URI")
	if len(dbUri) == 0 {
		return errors.New("env for dbUri is not set")
	}

	var err error
	db, err = tidb.Open(dbUri)
	return err
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

	orgId, err := db.ReadUserOrgIdEm(email)
	if err != nil {
		log.Warn(err)
		return response.Internal(), nil
	}

	enc := base64.StdEncoding.EncodeToString(orgId)

	return response.Text(enc), nil
}

func main() {
	log.FmtPlain()

	if err := connectDb(); err != nil {
		log.Fatal(err)
	}

	lambda.Start(handler)
}
