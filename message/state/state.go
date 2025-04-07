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

	ST_MCALL_ALC_PHOTO
	ST_MCALL_HEALTH
	ST_MCALL_CAR_CHECK
	ST_MCALL_NOTE

	ST_ECALL_ALC_PHOTO
	ST_ECALL_NOTE

	ST_DREPO_MTR_PHOTO
	ST_DREPO_TROUBLE
	ST_DREPO_NOTE

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
	case ST_MCALL_ALC_PHOTO:
		return StMCallAlcPhoto{}, nil
	case ST_MCALL_HEALTH:
		return StMCallHealth{}, nil
	case ST_MCALL_CAR_CHECK:
		return StMCallCarCheck{}, nil
	case ST_MCALL_NOTE:
		return StMCallNote{}, nil
	case ST_ECALL_ALC_PHOTO:
		return StECallAlcPhoto{}, nil
	case ST_ECALL_NOTE:
		return StECallNote{}, nil
	case ST_DREPO_MTR_PHOTO:
		return StDRepoMtrPhoto{}, nil
	case ST_DREPO_TROUBLE:
		return StDRepoTrouble{}, nil
	case ST_DREPO_NOTE:
		return StDRepoNote{}, nil

	case ST_IN_DUTY:
		return StInDuty{}, nil
	case ST_NONE:
		return StNone{}, nil

	default:
		return nil, errors.New("undefined state index")
	}
}
