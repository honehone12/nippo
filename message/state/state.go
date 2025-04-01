package state

import (
	"errors"
	"nippo/message/property"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type State interface {
	Run(
		props *property.MessagingProps,
		user *schema.User,
		event webhook.EventInterface,
	) error
}

func FromIdx(idx schema.StateIdx) (State, error) {
	switch idx {
	case ST_IDX_NONE:
		return &StNone{idx}, nil
	case ST_USER_NOT_CONFIGURED:
		return &StUserNotConfigured{idx}, nil
	default:
		return nil, errors.New("undefined state index")
	}
}
