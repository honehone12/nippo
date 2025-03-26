package response

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func BadRequest() *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       "bad request",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}
}

func Ok() *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "ok",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}
}
