package state

import (
	"errors"
	"nippo/message/msgprops"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

const (
	ST_USER_NOT_CONFIGURED schema.StateIdx = iota

	ST_USER_SET_NAME
	ST_USER_SET_EMAIL

	ST_MCALL_ALC_PHOTO
	ST_MCALL_HEALTH
	ST_MCALL_CAR_CHECK
	ST_MCALL_SET_CAR
	ST_MCALL_INSPECTION_EXP
	ST_MCALL_LICENSE_EXP
	ST_MCALL_NOTE

	ST_DREPO_M_MTR
	ST_DREPO_M_MTR_PHOTO

	ST_ECALL_ALC_PHOTO
	ST_ECALL_NOTE

	ST_DREPO_E_MTR
	ST_DREPO_E_MTR_PHOTO
	ST_DREPO_NON_DUTY_DST
	ST_DREPO_TROUBLE
	ST_DREPO_NOTE

	ST_SITE_NAME
	ST_SITE_NOTE

	ST_LOC_NOTE

	ST_REST_LOC
	ST_REST_NOTE

	ST_WAITING
	ST_WAIT_NOTE

	ST_LOADING
	ST_LOAD_SHIP_CHECK
	ST_LOAD_NOTE

	ST_IN_DUTY
	ST_NONE
)

type State interface {
	Run(
		props *msgprops.Props,
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
	case ST_USER_SET_EMAIL:
		return StUserSetEmail{}, nil
	case ST_MCALL_ALC_PHOTO:
		return StMCallAlcPhoto{}, nil
	case ST_MCALL_HEALTH:
		return StMCallHealth{}, nil
	case ST_MCALL_CAR_CHECK:
		return StMCallCarCheck{}, nil
	case ST_MCALL_SET_CAR:
		return StMCallSetCar{}, nil
	case ST_MCALL_INSPECTION_EXP:
		return StMCallInspectionExp{}, nil
	case ST_MCALL_LICENSE_EXP:
		return StMCallLicenseExp{}, nil
	case ST_MCALL_NOTE:
		return StMCallNote{}, nil
	case ST_DREPO_M_MTR:
		return StDRepoMMtr{}, nil
	case ST_DREPO_M_MTR_PHOTO:
		return StDRepoMMtrPhoto{}, nil
	case ST_ECALL_ALC_PHOTO:
		return StECallAlcPhoto{}, nil
	case ST_ECALL_NOTE:
		return StECallNote{}, nil
	case ST_DREPO_E_MTR:
		return StDRepoEMtr{}, nil
	case ST_DREPO_E_MTR_PHOTO:
		return StDRepoEMtrPhoto{}, nil
	case ST_DREPO_NON_DUTY_DST:
		return StDRepoNonDutyDst{}, nil
	case ST_DREPO_TROUBLE:
		return StDRepoTrouble{}, nil
	case ST_DREPO_NOTE:
		return StDRepoNote{}, nil
	case ST_SITE_NAME:
		return StSiteName{}, nil
	case ST_SITE_NOTE:
		return StSiteNote{}, nil
	case ST_LOC_NOTE:
		return StLocNote{}, nil
	case ST_REST_LOC:
		return StResting{}, nil
	case ST_REST_NOTE:
		return StRestNote{}, nil
	case ST_WAITING:
		return StWaiting{}, nil
	case ST_WAIT_NOTE:
		return StWaitNote{}, nil
	case ST_LOADING:
		return StLoading{}, nil
	case ST_LOAD_SHIP_CHECK:
		return StLoadShipCheck{}, nil
	case ST_LOAD_NOTE:
		return StLoadNote{}, nil

	case ST_IN_DUTY:
		return StInDuty{}, nil
	case ST_NONE:
		return StNone{}, nil

	default:
		return nil, errors.New("undefined state index")
	}
}
