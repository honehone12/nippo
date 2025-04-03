package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

const USER_SETTING_ACT_DATA = "USER_SETTING"
const START_DUTY_ACT_DATA = "START_DUTY"
const FLAG_TRUE_ACT_DATA = "FLAG_TRUE"

func ActionUserSetting() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "ユーザー設定を行う",
		Data:        USER_SETTING_ACT_DATA,
		DisplayText: "ユーザー設定を行う",
		InputOption: "openKeyboard",
	}
}

func ActionSendAlcPhoto() messaging_api.CameraAction {
	return messaging_api.CameraAction{
		Label: "アルコール検査機の結果を写真に撮る",
	}
}

func ActionFlagTrue() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "良好",
		Data:        FLAG_TRUE_ACT_DATA,
		DisplayText: "良好",
	}
}

func ActionStartDuty() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "業務を開始する",
		Data:        START_DUTY_ACT_DATA,
		DisplayText: "業務を開始する",
	}
}
