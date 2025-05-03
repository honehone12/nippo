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
	"nippo/message/msgprops"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type variable struct {
	lnChSec string
	orgId   uuid.UUID
	bucket  string
}

var (
	env      variable
	db       database.Database
	api      *messaging_api.MessagingApiAPI
	blobApi  *messaging_api.MessagingApiBlobAPI
	s3Client *s3.Client
)

func cacheEnv() error {
	if env.lnChSec = os.Getenv("LN_CH_SEC"); len(env.lnChSec) == 0 {
		return errors.New("env for lnChSec is not set")
	}

	orgIdStr := os.Getenv("ORG_ID")
	if len(orgIdStr) == 0 {
		return errors.New("env for orgIdStr is not set")
	}
	orgId, err := uuid.Parse(orgIdStr)
	if err != nil {
		return err
	}
	env.orgId = orgId

	if env.bucket = os.Getenv("BUCKET"); len(env.bucket) == 0 {
		return errors.New("env for bucket is not set")
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

func initApi() error {
	lnChTkn := os.Getenv("LN_CH_TKN")
	if len(lnChTkn) == 0 {
		return errors.New("env for lnChTkn is not set")
	}

	var err error
	api, err = messaging_api.NewMessagingApiAPI(lnChTkn)
	if err != nil {
		return err
	}

	blobApi, err = messaging_api.NewMessagingApiBlobAPI(lnChTkn)
	if err != nil {
		return err
	}

	return nil
}

func configAws() error {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return err
	}

	s3Client = s3.NewFromConfig(cfg)
	return nil
}

func validateHeaders(req *events.APIGatewayProxyRequest) error {
	userAgent, ok := req.Headers["User-Agent"]
	if !ok {
		userAgent, ok = req.Headers["user-agent"]
	}
	if !ok || !strings.Contains(userAgent, "LineBotWebhook") {
		return errors.New("invalid user agent")
	} else if userAgent != "LineBotWebhook/2.0" {
		log.Warn("user agent was updated from version2.0")
	}

	contentType, ok := req.Headers["Content-Type"]
	if !ok {
		contentType, ok = req.Headers["content-type"]
	}
	if !ok || !strings.Contains(contentType, "application/json") {
		return errors.New("invalid content type header")
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

	sig, ok := req.Headers["x-line-signature"]
	if !ok {
		return errors.New("no signature")
	}
	if !webhook.ValidateSignature(env.lnChSec, sig, b) {
		return errors.New("invalid signature")
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
	if err := validateHeaders(&req); err != nil {
		log.Warn(err)
		return response.BadRequest(), nil
	}

	var msg webhook.CallbackRequest
	if err := parseMsg(&req, &msg); err != nil {
		log.Warn(err)
		return response.BadRequest(), nil
	}

	props := msgprops.NewProps(
		ctx,
		env.bucket,
		env.orgId,
		db,
		api,
		blobApi,
		s3Client,
	)
	message.MessageMain(props, &msg)

	return response.Text("ok"), nil
}

func main() {
	log.FmtPlain()

	if err := cacheEnv(); err != nil {
		log.Fatal(err)
	}

	if err := configAws(); err != nil {
		log.Fatal(err)
	}

	if err := initApi(); err != nil {
		log.Fatal(err)
	}

	if err := connectDb(); err != nil {
		log.Fatal(err)
	}

	lambda.Start(handler)
}
