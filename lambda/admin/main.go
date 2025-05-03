package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"net/url"
	"nippo/admin"
	"nippo/admin/adminprops"
	"nippo/database"
	"nippo/database/tidb"
	"nippo/input"
	"nippo/lambda/response"
	"nippo/log"
	"nippo/schema"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/google/uuid"
)

type variable struct {
	orgId    uuid.UUID
	userPool string
}

var env variable
var db database.Database
var cognitoClient *cognitoidentityprovider.Client

func cacheEnv() error {
	orgIdStr := os.Getenv("ORG_ID")
	if len(orgIdStr) == 0 {
		return errors.New("env for orgIdStr is not set")
	}
	orgId, err := uuid.Parse(orgIdStr)
	if err != nil {
		return err
	}
	env.orgId = orgId

	if env.userPool = os.Getenv("USER_POOL"); len(env.userPool) == 0 {
		return errors.New("env for user pool is not set")
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

func configAws() error {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return err
	}

	cognitoClient = cognitoidentityprovider.NewFromConfig(cfg)

	return nil
}

func validateHeaders(headers map[string]string) error {
	contentType, ok := headers["Content-Type"]
	if !ok {
		contentType, ok = headers["content-type"]
	}
	if !ok || !strings.Contains(contentType, "application/x-www-form-urlencoded") {
		return errors.New("invalid content type header")
	}

	return nil
}

func adminEmail(auth map[string]any) (schema.CleanText, error) {
	claims, ok := auth["claims"].(map[string]any)
	if !ok {
		return "", errors.New("auth not found")
	}

	raw, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("user email not found")
	}

	email, err := input.ValidEmail(raw)
	if err != nil {
		return "", err
	}

	return email, nil
}

func parse(body string, isBase64 bool) (url.Values, error) {
	var s string
	if isBase64 {
		b, err := base64.StdEncoding.DecodeString(body)
		if err != nil {
			return nil, err
		}
		s = string(b)
	} else {
		s = body
	}

	return url.ParseQuery(s)
}

func handler(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	if err := validateHeaders(req.Headers); err != nil {
		log.Warn(err)
		return response.BadRequest(), nil
	}

	adminEmail, err := adminEmail(req.RequestContext.Authorizer)
	if err != nil {
		log.Warn(err)
		return response.BadRequest(), nil
	}

	adminOrg, err := db.ReadUserOrgId_Email(adminEmail)
	if err != nil {
		log.Warn(err)
		return response.Internal(), nil
	}
	if !bytes.Equal(adminOrg, env.orgId[:]) {
		log.Warn("viewing user has wrong org id")
		return response.BadRequest(), nil
	}

	orgEmail, err := db.ReadOrgEmail(adminOrg)
	if err != nil {
		log.Warn(err)
		return response.Internal(), nil
	}
	if orgEmail != adminEmail {
		log.Warn("not admin")
		return response.BadRequest(), nil
	}

	form, err := parse(req.Body, req.IsBase64Encoded)
	if err != nil {
		log.Warn(err)
		return response.BadRequest(), nil
	}

	props := adminprops.NewProps(
		ctx,
		env.userPool,
		env.orgId,
		db,
		cognitoClient,
	)

	switch form.Get("op") {
	case "invite":
		email, err := admin.Invite(props, form)
		if err != nil {
			return response.Internal(), nil
		}
		return response.Text(string(email)), nil
	default:
		log.Warn("invalid operation")
		return response.BadRequest(), nil
	}
}

func main() {
	log.FmtPlain()

	if err := cacheEnv(); err != nil {
		log.Fatal(err)
	}

	if err := configAws(); err != nil {
		log.Fatal(err)
	}

	if err := connectDb(); err != nil {
		log.Fatal(err)
	}

	lambda.Start(handler)
}
