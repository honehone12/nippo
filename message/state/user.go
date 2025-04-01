package state

import (
	"nippo/message/property"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StUserNotConfigured struct {
	idx schema.StateIdx
}

func (s *StUserNotConfigured) Run(
	props *property.MsgProps,
	event webhook.EventInterface,
) error {
	panic("not implemented")
}

func UserNotConfigured() State {
	return &StUserNotConfigured{idx: schema.ST_USER_NOT_CONFIGURED}
}
