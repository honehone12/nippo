package msgprops

import (
	"context"
	"nippo/database"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type Props struct {
	ctx     context.Context
	bucket  string
	orgId   uuid.UUID
	db      database.Database
	api     *messaging_api.MessagingApiAPI
	blobApi *messaging_api.MessagingApiBlobAPI
	s3      *s3.Client
}

func NewProps(
	ctx context.Context,
	bucket string,
	orgId uuid.UUID,
	db database.Database,
	api *messaging_api.MessagingApiAPI,
	blobApi *messaging_api.MessagingApiBlobAPI,
	s3 *s3.Client,
) *Props {
	return &Props{ctx, bucket, orgId, db, api, blobApi, s3}
}

func (p *Props) Ctx() context.Context {
	return p.ctx
}

func (p *Props) Bucket() *string {
	return &p.bucket
}

func (p *Props) OrgId() []byte {
	return p.orgId[:]
}

func (p *Props) OrgUuid() uuid.UUID {
	return p.orgId
}

func (p *Props) Db() database.Database {
	return p.db
}

func (p *Props) Api() *messaging_api.MessagingApiAPI {
	return p.api
}

func (p *Props) BlobApi() *messaging_api.MessagingApiBlobAPI {
	return p.blobApi
}

func (p *Props) S3() *s3.Client {
	return p.s3
}
