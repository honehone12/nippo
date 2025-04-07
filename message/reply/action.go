package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

const IGNORE_TXT_TEMPLATE = "IGNORE_TXT_TEMPLATE"
const USER_SETTING_ACT_DATA = "USER_SETTING"
const START_DUTY_ACT_DATA = "START_DUTY"
const START_WAITING_ACT_DATA = "START_WAITING"
const START_LOADING_ACT_DATA = "START_LOADING"
const START_RESTING_ACT_DATA = "START_RESTING"
const END_DUTY_ACT_DATA = "END_DUTY"
const FLAG_TRUE_ACT_DATA = "FLAG_TRUE"
const RESET_ACT_DATA = "RESET"

func ActionUserSetting() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "ユーザー設定を行う",
		Data:        USER_SETTING_ACT_DATA,
		DisplayText: "ユーザー設定を行う",
		InputOption: "openKeyboard",
	}
}

func ActionMCallNote() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "入力する",
		Data:        IGNORE_TXT_TEMPLATE,
		InputOption: "openKeyboard",
		FillInText:  "【指示事項】\n\n【天気】\n",
	}
}

func ActionECallNote() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "入力する",
		Data:        IGNORE_TXT_TEMPLATE,
		InputOption: "openKeyboard",
		FillInText:  "【指示事項】\n\n【交替運転者に対する通告】\n\n【車両・道路及び運行の状況】\n\n【その他】\n",
	}
}

func ActionSendAlcPhoto() messaging_api.CameraAction {
	return messaging_api.CameraAction{
		Label: "アルコール検査機の結果を写真に撮る",
	}
}

func ActionSendMtrPhoto() messaging_api.CameraAction {
	return messaging_api.CameraAction{
		Label: "車両のメーターを写真に撮る",
	}
}

func ActionDReportTrouble() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "入力する",
		Data:        IGNORE_TXT_TEMPLATE,
		InputOption: "openKeyboard",
		FillInText:  "【事故】\n\n【著しい遅延】\n\n【その他の異常】\n",
	}
}

func ActionDReportNote() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "入力する",
		Data:        IGNORE_TXT_TEMPLATE,
		InputOption: "openKeyboard",
		FillInText:  "【業務の交替がある場合】\n地点：　\n 時　分\n【高速利用】\n\n【その他】\n",
	}
}

func ActionSendLocation() messaging_api.LocationAction {
	return messaging_api.LocationAction{
		Label: "位置情報を登録する",
	}
}

func ActionLocNote() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "入力する",
		Data:        IGNORE_TXT_TEMPLATE,
		InputOption: "openKeyboard",
		FillInText:  "【説明】\n",
	}
}

func ActionStartWaiting() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "待機を開始する",
		Data:        START_WAITING_ACT_DATA,
		DisplayText: "待機を開始する",
	}
}

func ActionStartLoading() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "荷役作業を開始する",
		Data:        START_LOADING_ACT_DATA,
		DisplayText: "荷役作業を開始する",
	}
}

func ActionStartResting() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "休憩を開始する",
		Data:        START_RESTING_ACT_DATA,
		DisplayText: "休憩を開始する",
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

func ActionEndDuty() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "業務を終了する",
		Data:        END_DUTY_ACT_DATA,
		DisplayText: "業務を終了する",
	}
}

func ActionReset() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "リセット",
		Data:        RESET_ACT_DATA,
		DisplayText: "リセット",
	}
}
