package state

import (
	"nippo/handlerprops"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type State interface {
	run(
		props *handlerprops.MessagingProperties,
		content webhook.MessageContentInterface,
	) error
}
