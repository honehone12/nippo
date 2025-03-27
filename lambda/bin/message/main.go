package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"nippo/lambda/message"
	"nippo/lambda/response"
	"nippo/log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type variables struct {
	lnChSec string
}

var (
	env variables
)

func collectEnv() error {
	env.lnChSec = os.Getenv("LN_CH_SEC")
	if len(env.lnChSec) == 0 {
		return errors.New("env for lnchsec is not set")
	}

	return nil
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
			return errors.New("failed to decode base64")
		}
	} else {
		b = []byte(req.Body)
	}

	if err := json.Unmarshal(b, msg); err != nil {
		return errors.New("failed to unmarshal json body")
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

	if err := message.MessageMain(&msg); err != nil {
		log.L.Error(err)
	}
	return *response.Ok(), nil
}

func main() {
	log.FmtDefault()

	if err := collectEnv(); err != nil {
		panic(err)
	}

	lambda.Start(handler)
}
