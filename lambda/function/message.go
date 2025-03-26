package function

import (
	"nippo/lambda/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func MessageMain(msg *webhook.CallbackRequest) *events.APIGatewayProxyResponse {

	return response.Ok()
}
