package state

import (
	"nippo/message"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type State interface {
	run(
		props *message.MessagingProperties,
		content webhook.MessageContentInterface,
	) error
}
