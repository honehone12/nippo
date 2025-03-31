package handlerprops

import (
	"nippo/database"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type MessagingProperties struct {
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
) *MessagingProperties {
	return &MessagingProperties{
		orgIdStr,
		db,
		api,
		blobApi,
	}
}
