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
	ST_MC_CAR_CHECK
	ST_MC_NOTE

	ST_EC_ALC_PHOTO
	ST_EC_NOTE

	ST_DR_MTR_PHOTO

	ST_IN_DUTY
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
	case ST_MC_HEALTH:
		return StMcHealth{}, nil
	case ST_MC_CAR_CHECK:
		return StMcCarCheck{}, nil
	case ST_MC_NOTE:
		return StMcNote{}, nil
	case ST_EC_ALC_PHOTO:
		return StEcAlcPhoto{}, nil
	case ST_EC_NOTE:
		return StEcNote{}, nil
	case ST_DR_MTR_PHOTO:
		return StDrMtrPhoto{}, nil

	case ST_IN_DUTY:
		return StInDuty{}, nil
	case ST_NONE:
		return StNone{}, nil

	default:
		return nil, errors.New("undefined state index")
	}
}
