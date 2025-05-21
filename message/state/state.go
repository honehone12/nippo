package state

import (
	"errors"
	"nippo/message/msgprops"
	"nippo/message/state/drepo"
	"nippo/message/state/ecall"
	"nippo/message/state/index"
	"nippo/message/state/loading"
	"nippo/message/state/location"
	"nippo/message/state/mcall"
	"nippo/message/state/resting"
	"nippo/message/state/site"
	"nippo/message/state/usrset"
	"nippo/message/state/waiting"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
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
	case index.USR_NOT_CONFIGURED:
		return UsrNotConfigured{}, nil

	case index.USRSET_NAME:
		return usrset.UsrsetName{}, nil
	case index.USRSET_EMAIL:
		return usrset.UsrsetEmail{}, nil

	case index.MCALL_ALC_PHOTO:
		return mcall.MCallAlcPhoto{}, nil
	case index.MCALL_HEALTH:
		return mcall.MCallHealth{}, nil
	case index.MCALL_CAR_CHECK:
		return mcall.MCallCarCheck{}, nil
	case index.MCALL_CAR:
		return mcall.MCallCar{}, nil
	case index.MCALL_INSPECTION_EXP:
		return mcall.MCallInspectionExp{}, nil
	case index.MCALL_LICENSE_EXP:
		return mcall.MCallLicenseExp{}, nil
	case index.MCALL_CALLER:
		return mcall.MCallCaller{}, nil
	case index.MCALL_NOTE:
		return mcall.MCallNote{}, nil

	case index.DREPO_M_MTR:
		return drepo.DRepoMMtr{}, nil
	case index.DREPO_M_MTR_PHOTO:
		return drepo.DRepoMMtrPhoto{}, nil

	case index.ECALL_CALLER:
		return ecall.ECallCaller{}, nil
	case index.ECALL_ALC_PHOTO:
		return ecall.ECallAlcPhoto{}, nil
	case index.ECALL_NOTE:
		return ecall.ECallNote{}, nil

	case index.DREPO_E_MTR:
		return drepo.DRepoEMtr{}, nil
	case index.DREPO_E_MTR_PHOTO:
		return drepo.DRepoEMtrPhoto{}, nil
	case index.DREPO_NON_DUTY_DST:
		return drepo.DRepoNonDutyDst{}, nil
	case index.DREPO_TROUBLE:
		return drepo.DRepoTrouble{}, nil
	case index.DREPO_NOTE:
		return drepo.DRepoNote{}, nil

	case index.SITE_NAME:
		return site.SiteName{}, nil
	case index.SITE_NOTE:
		return site.SiteNote{}, nil

	case index.LOC_NOTE:
		return location.LocNote{}, nil

	case index.REST_LOC:
		return resting.Resting{}, nil
	case index.REST_NOTE:
		return resting.RestNote{}, nil

	case index.WAITING:
		return waiting.Waiting{}, nil
	case index.WAIT_NOTE:
		return waiting.WaitNote{}, nil

	case index.LOADING:
		return loading.Loading{}, nil
	case index.LOAD_SHIP_CHECK:
		return loading.LoadShipCheck{}, nil
	case index.LOAD_NOTE:
		return loading.LoadNote{}, nil

	case index.IN_DUTY:
		return StInDuty{}, nil
	case index.NONE:
		return StNone{}, nil

	default:
		return nil, errors.New("undefined state index")
	}
}
