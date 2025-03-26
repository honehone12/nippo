package request

import (
	"encoding/base64"
	"encoding/json"
	"nippo/lambda/response"
	"nippo/log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func CheckHeader(
	req *events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, bool) {
	if h, ok := req.Headers["Content-Type"]; !ok || h != "application/json" {
		log.L.Warn("invalid content type header")
		return response.BadRequest(), false
	}
	sig, ok := req.Headers["x-line-signature"]
	if !ok {
		log.L.Warn("no signature")
		return response.BadRequest(), false
	}
	lineSec := os.Getenv("LINE_CHANNEL_SECRET")
	if len(lineSec) == 0 {
		log.L.Error("env LINE_CHANNEL_SECRET is not set")
		return response.Ok(), false
	}
	if !webhook.ValidateSignature(lineSec, sig, []byte(req.Body)) {
		log.L.Warn("invalid signature")
		return response.BadRequest(), false
	}

	return nil, true
}

func Parse(
	req *events.APIGatewayProxyRequest,
	msg *webhook.CallbackRequest,
) (*events.APIGatewayProxyResponse, bool) {
	var b []byte
	if req.IsBase64Encoded {
		var err error
		b, err = base64.StdEncoding.DecodeString(req.Body)
		if err != nil {
			log.L.Error("failed to decode base64")
			return response.Ok(), false
		}
	} else {
		b = []byte(req.Body)
	}

	if err := json.Unmarshal(b, msg); err != nil {
		log.L.Warn("failed to unmarshal json body")
		return response.BadRequest(), false
	}

	return nil, true
}
