package state

import (
	"nippo/message/property"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StNone struct {
	idx schema.StateIdx
}

func (s *StNone) Run(
	props *property.MsgProps,
	event webhook.EventInterface,
) error {
	panic("not implemented")
}
