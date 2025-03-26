package main

import (
	"context"
	"nippo/lambda/function"
	"nippo/lambda/request"
	"nippo/log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func handler(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	if res, ok := request.CheckHeader(&req); !ok {
		return *res, nil
	}

	var msg webhook.CallbackRequest
	if res, ok := request.Parse(&req, &msg); !ok {
		return *res, nil
	}

	return *function.MessageMain(&msg), nil
}

func init() {
	log.FmtDefault()
}

func main() {
	lambda.Start(handler)
}
