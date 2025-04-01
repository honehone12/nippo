package action

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

const USER_SETTING_ACTION_DATA = "USER_SETTING"

func UserSettingAction() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "ユーザー設定を行う",
		Data:        USER_SETTING_ACTION_DATA,
		DisplayText: "ユーザー設定を行う",
		InputOption: "openKeyboard",
	}
}
