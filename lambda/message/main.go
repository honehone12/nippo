package main

import (
	"context"
	"encoding/base64"
	"errors"
	"nippo/database"
	"nippo/database/tidb"
	"nippo/lambda/response"
	"nippo/log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type variable struct {
	lnChSec  string
	orgIdStr string
	dbUri    string
}

var (
	env variable
	db  database.Database
)

func collectEnv() error {
	if env.lnChSec = os.Getenv("LN_CH_SEC"); len(env.lnChSec) == 0 {
		return errors.New("env for lnChSec is not set")
	}

	if env.orgIdStr = os.Getenv("ORG_ID"); len(env.orgIdStr) == 0 {
		return errors.New("env for orgIdStr is not set")
	}

	if env.dbUri = os.Getenv("DB_URI"); len(env.dbUri) == 0 {
		return errors.New("env for dbUri is not set")
	}

	return nil
}

func connectDb() error {
	var err error
	db, err = tidb.Open(env.dbUri)
	return err
}

func validateMsg(req *events.APIGatewayProxyRequest) error {
	if h, ok := req.Headers["Content-Type"]; !ok || h != "application/json" {
		return errors.New("invalid content type header")
	}
	sig, ok := req.Headers["x-line-signature"]
	if !ok {
		return errors.New("no signature")
	}

	if !webhook.ValidateSignature(env.lnChSec, sig, []byte(req.Body)) {
		return errors.New("invalid signature")
	}

	return nil
}

func parseMsg(
	req *events.APIGatewayProxyRequest,
	msg *webhook.CallbackRequest,
) error {
	var b []byte
	if req.IsBase64Encoded {
		var err error
		b, err = base64.StdEncoding.DecodeString(req.Body)
		if err != nil {
			return err
		}
	} else {
		b = []byte(req.Body)
	}

	if err := msg.UnmarshalJSON(b); err != nil {
		return err
	}

	return nil
}

func handler(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	if err := validateMsg(&req); err != nil {
		log.L.Warn(err)
		return *response.BadRequest(), nil
	}

	var msg webhook.CallbackRequest
	if err := parseMsg(&req, &msg); err != nil {
		log.L.Warn(err)
		return *response.BadRequest(), nil
	}

	log.L.Infof("%#v\n", msg)

	// if err := message.MessageMain(&message.MessageProps{
	// 	OrgIdStr: env.orgIdStr,
	// 	Db:       db,
	// }, &msg); err != nil {
	// 	log.L.Error(err)
	// }
	return *response.Ok(), nil
}

func main() {
	log.FmtDefault()

	if err := collectEnv(); err != nil {
		log.L.Fatal(err)
	}

	lambda.Start(handler)
}
