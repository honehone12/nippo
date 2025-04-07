package property

import (
	"context"
	"nippo/database"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type MsgingProps struct {
	ctx      context.Context
	bucket   string
	orgIdStr string
	db       database.Database
	api      *messaging_api.MessagingApiAPI
	blobApi  *messaging_api.MessagingApiBlobAPI
	s3       *s3.Client
}

func NewMessagingProps(
	ctx context.Context,
	bucket string,
	orgIdStr string,
	db database.Database,
	api *messaging_api.MessagingApiAPI,
	blobApi *messaging_api.MessagingApiBlobAPI,
	s3 *s3.Client,
) *MsgingProps {
	return &MsgingProps{
		ctx,
		bucket,
		orgIdStr,
		db,
		api,
		blobApi,
		s3,
	}
}

func (p *MsgingProps) Ctx() context.Context {
	return p.ctx
}

func (p *MsgingProps) Bucket() *string {
	return &p.bucket
}

func (p *MsgingProps) OrgId() string {
	return p.orgIdStr
}

func (p *MsgingProps) Db() database.Database {
	return p.db
}

func (p *MsgingProps) Api() *messaging_api.MessagingApiAPI {
	return p.api
}

func (p *MsgingProps) BlobApi() *messaging_api.MessagingApiBlobAPI {
	return p.blobApi
}

func (p *MsgingProps) S3() *s3.Client {
	return p.s3
}
