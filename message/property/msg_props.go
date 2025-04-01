package property

import (
	"nippo/database"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type MsgProps struct {
	orgIdStr string
	db       database.Database
	api      *messaging_api.MessagingApiAPI
	blobApi  *messaging_api.MessagingApiBlobAPI
}

func NewMessageProps(
	orgIdStr string,
	db database.Database,
	api *messaging_api.MessagingApiAPI,
	blobApi *messaging_api.MessagingApiBlobAPI,
) *MsgProps {
	return &MsgProps{
		orgIdStr,
		db,
		api,
		blobApi,
	}
}

func (p *MsgProps) OrgId() string {
	return p.orgIdStr
}

func (p *MsgProps) Db() database.Database {
	return p.db
}

func (p *MsgProps) Api() *messaging_api.MessagingApiAPI {
	return p.api
}

func (p *MsgProps) BlobApi() *messaging_api.MessagingApiBlobAPI {
	return p.blobApi
}
