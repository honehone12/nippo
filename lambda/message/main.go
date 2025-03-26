package main

import (
	"context"
	"net/http"
	"nippo/log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	log.L.Infof("%t:%s\n", req.IsBase64Encoded, req.Body)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "ok",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}, nil
}

func init() {
	log.FmtDefault()
}

func main() {
	lambda.Start(handler)
}
