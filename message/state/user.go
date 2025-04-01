package state

import (
	"nippo/message/property"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StUseNotConfigured struct {
	idx schema.StateIdx
}

func (s *StUseNotConfigured) Run(
	props *property.MessagingProperties,
	event webhook.EventInterface,
) error {
	panic("not implemented")
}
