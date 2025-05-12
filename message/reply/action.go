package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

const IGNORE_TXT_ACT_DATA = "IGNORE_TXT"
const USER_SETTING_ACT_DATA = "USER_SETTING"
const START_DUTY_ACT_DATA = "START_DUTY"
const START_WAITING_ACT_DATA = "START_WAITING"
const START_LOADING_ACT_DATA = "START_LOADING"
const REST_LOC_ACT_DATA = "START_RESTING"
const END_DUTY_ACT_DATA = "END_DUTY"
const FLAG_TRUE_ACT_DATA = "FLAG_TRUE"
const FLAG_FALSE_ACT_DATA = "FLAG_FALSE"
const DECLINE_ACT_DATA = "RESET"

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
		Data:        IGNORE_TXT_ACT_DATA,
		InputOption: "openKeyboard",
		FillInText:  "【指示事項】\n【天気】\n【健康状態】\n【車両状態】\n【その他】",
	}
}

func ActionECallNote() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "入力する",
		Data:        IGNORE_TXT_ACT_DATA,
		InputOption: "openKeyboard",
		FillInText:  "【指示事項】\n【交替運転者に対する通告】\n【車両・道路及び運行の状況】\n【その他】",
	}
}

func ActionAlcPhoto() messaging_api.CameraAction {
	return messaging_api.CameraAction{
		Label: "アルコール検査機の結果を写真に撮る",
	}
}

func ActionMtrPhoto() messaging_api.CameraAction {
	return messaging_api.CameraAction{
		Label: "車両のメーターを写真に撮る",
	}
}

func ActionDReportTrouble() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "入力する",
		Data:        IGNORE_TXT_ACT_DATA,
		InputOption: "openKeyboard",
		FillInText:  "【事故】\n【著しい遅延】\n【その他の異常】",
	}
}

func ActionDReportNote() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "入力する",
		Data:        IGNORE_TXT_ACT_DATA,
		InputOption: "openKeyboard",
		FillInText:  "【業務交替】時間：　場所：　\n【件数】集荷：　配達：　\n【高速利用】\n【その他】",
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
		Data:        IGNORE_TXT_ACT_DATA,
		InputOption: "openKeyboard",
		FillInText:  "【説明】",
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
		Label:       "荷役（附帯）作業を開始する",
		Data:        START_LOADING_ACT_DATA,
		DisplayText: "荷役（附帯）作業を開始する",
	}
}

func ActionEndLoading() messaging_api.LocationAction {
	return messaging_api.LocationAction{
		Label: "作業を終了する（位置情報を送信）",
	}
}

func ActionLoadNote() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "入力する",
		Data:        IGNORE_TXT_ACT_DATA,
		InputOption: "openKeyboard",
		FillInText:  "【作業内容】\n【時間】\n【その他】",
	}
}

func ActionRegisterResting() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "休憩を登録する",
		Data:        REST_LOC_ACT_DATA,
		DisplayText: "休憩を登録する",
	}
}

func ActionRestLoc() messaging_api.LocationAction {
	return messaging_api.LocationAction{
		Label: "位置情報を送信する",
	}
}

func ActionRestNote() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "入力する",
		Data:        IGNORE_TXT_ACT_DATA,
		InputOption: "openKeyboard",
		FillInText:  "【時間】\n【睡眠】",
	}
}

func ActionEndWaiting() messaging_api.LocationAction {
	return messaging_api.LocationAction{
		Label: "待機を終了する（位置情報を送信）",
	}
}

func ActionWaitNote() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "入力する",
		Data:        IGNORE_TXT_ACT_DATA,
		InputOption: "openKeyboard",
		FillInText:  "【荷主指定時刻】\n【時間】\n【その他】",
	}
}

func ActionFlagTrue() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "良い",
		Data:        FLAG_TRUE_ACT_DATA,
		DisplayText: "良い",
	}
}

func ActionFlagFalse() messaging_api.PostbackAction {
	return messaging_api.PostbackAction{
		Label:       "不明",
		Data:        FLAG_FALSE_ACT_DATA,
		DisplayText: "不明",
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
		Label:       "辞退",
		Data:        DECLINE_ACT_DATA,
		DisplayText: "辞退",
	}
}

func ActionNone() messaging_api.MessageAction {
	return messaging_api.MessageAction{
		Label: "無し",
		Text:  "無し",
	}
}
