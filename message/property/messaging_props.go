package property

import (
	"nippo/database"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type MessagingProps struct {
	orgIdStr string
	db       database.Database
	api      *messaging_api.MessagingApiAPI
	blobApi  *messaging_api.MessagingApiBlobAPI
}

func NewMessagingProps(
	orgIdStr string,
	db database.Database,
	api *messaging_api.MessagingApiAPI,
	blobApi *messaging_api.MessagingApiBlobAPI,
) *MessagingProps {
	return &MessagingProps{
		orgIdStr,
		db,
		api,
		blobApi,
	}
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
