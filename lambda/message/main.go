package main

import (
	"context"
	"encoding/base64"
	"errors"
	"nippo/database"
	"nippo/database/tidb"
	"nippo/lambda/response"
	"nippo/log"
	"nippo/message"
	"nippo/message/property"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type variable struct {
	lnChSec  string
	lnChTkn  string
	orgIdStr string
	dbUri    string
	bucket   string
}

var (
	env      variable
	db       database.Database
	api      *messaging_api.MessagingApiAPI
	blobApi  *messaging_api.MessagingApiBlobAPI
	s3Client *s3.Client
)

func collectEnv() error {
	if env.lnChSec = os.Getenv("LN_CH_SEC"); len(env.lnChSec) == 0 {
		return errors.New("env for lnChSec is not set")
	}

	if env.lnChTkn = os.Getenv("LN_CH_TKN"); len(env.lnChTkn) == 0 {
		return errors.New("env for lnChTkn is not set")
	}

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

func initApi() error {
	var err error
	api, err = messaging_api.NewMessagingApiAPI(env.lnChTkn)
	if err != nil {
		return err
	}

	blobApi, err = messaging_api.NewMessagingApiBlobAPI(env.lnChTkn)
	if err != nil {
		return err
	}

	return nil
}

func awsConfig() error {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return err
	}

	s3Client = s3.NewFromConfig(cfg)
	return nil
}

func validateMsg(req *events.APIGatewayProxyRequest) error {
	userAgent, ok := req.Headers["User-Agent"]
	if !ok {
		userAgent, ok = req.Headers["user-agent"]
	}
	if !ok || !strings.Contains(userAgent, "LineBotWebhook") {
		return errors.New("invalid user agent")
	} else if strings.Compare(userAgent, "LineBotWebhook/2.0") != 0 {
		log.L.Warn("user agent was updated from version2.0")
	}

	contentType, ok := req.Headers["Content-Type"]
	if !ok {
		contentType, ok = req.Headers["content-type"]
	}
	if !ok || !strings.Contains(contentType, "application/json") {
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
		return response.BadRequest(), nil
	}

	var msg webhook.CallbackRequest
	if err := parseMsg(&req, &msg); err != nil {
		log.L.Warn(err)
		return response.BadRequest(), nil
	}

	props := property.NewMessagingProps(
		ctx,
		env.bucket,
		env.orgIdStr,
		db,
		api,
		blobApi,
		s3Client,
	)
	message.MessageMain(props, &msg)

	return response.Ok(), nil
}

func main() {
	log.FmtDefault()

	if err := collectEnv(); err != nil {
		log.L.Fatal(err)
	}

	if err := connectDb(); err != nil {
		log.L.Fatal(err)
	}

	if err := initApi(); err != nil {
		log.L.Fatal(err)
	}

	if err := awsConfig(); err != nil {
		log.L.Fatal(err)
	}

	lambda.Start(handler)
}
