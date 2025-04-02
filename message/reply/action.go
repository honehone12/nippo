package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

const USER_SETTING_ACT_DATA = "USER_SETTING"
const START_DUTY_ACT_DATA = "START_DUTY"

func UserSettingAction() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "ユーザー設定を行う",
		Data:        USER_SETTING_ACT_DATA,
		DisplayText: "ユーザー設定を行う",
		InputOption: "openKeyboard",
	}
}

func StartDutyAction() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "業務を開始する",
		Data:        START_DUTY_ACT_DATA,
		DisplayText: "業務を開始する",
	}
}

func SendAlcPhotoAction() messaging_api.CameraAction {
	return messaging_api.CameraAction{
		Label: "アルコール検査機の結果を写真に撮る",
	}
}
