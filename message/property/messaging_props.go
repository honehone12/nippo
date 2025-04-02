package property

import (
	"context"
	"nippo/database"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type MessagingProps struct {
	ctx      context.Context
	orgIdStr string
	db       database.Database
	api      *messaging_api.MessagingApiAPI
	blobApi  *messaging_api.MessagingApiBlobAPI
}

func NewMessagingProps(
	ctx context.Context,
	orgIdStr string,
	db database.Database,
	api *messaging_api.MessagingApiAPI,
	blobApi *messaging_api.MessagingApiBlobAPI,
) *MessagingProps {
	return &MessagingProps{
		ctx,
		orgIdStr,
		db,
		api,
		blobApi,
	}
}

func (p *MessagingProps) Ctx() context.Context {
	return p.ctx
}

func (p *MessagingProps) OrgId() string {
	return p.orgIdStr
}

func (p *MessagingProps) Db() database.Database {
	return p.db
}

func (p *MessagingProps) Api() *messaging_api.MessagingApiAPI {
	return p.api
}

func (p *MessagingProps) BlobApi() *messaging_api.MessagingApiBlobAPI {
	return p.blobApi
}
