package state

import (
	"nippo/message/property"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StDrNote struct{}

func (s StDrNote) Run(
	props *property.MessagingProps,
	user *schema.User,
	event webhook.EventInterface,
) error {
	panic("not implemented")
}
