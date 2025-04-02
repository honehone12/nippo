package state

import (
	"errors"
	"nippo/message/property"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

const (
	ST_USER_NOT_CONFIGURED schema.StateIdx = iota

	ST_USER_SET_NAME
	ST_USER_SET_CAR

	ST_MC_ALC_PHOTO
	ST_MC_HEALTH

	ST_NONE
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
	case ST_USER_NOT_CONFIGURED:
		return StUserNotConfigured{}, nil
	case ST_USER_SET_NAME:
		return StUserSetName{}, nil
	case ST_USER_SET_CAR:
		return StUserSetCar{}, nil
	case ST_MC_ALC_PHOTO:
		return StMcAlcPhoto{}, nil
	case ST_NONE:
		return StNone{}, nil
	default:
		return nil, errors.New("undefined state index")
	}
}
