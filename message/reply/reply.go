package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

func ReplyDoUsrSet() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "ユーザー設定を行ってください",
			QuickReply: QuickRepConfigMenu(),
		},
	}
}

func ReplyUsrSetName() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "名前を入力してください",
		},
	}
}

func ReplyUsrSetCarNum() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "車両番号を入力してください",
		},
	}
}

func ReplySendAlcPhoto() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "アルコール検査機を使用し、結果が良好な場合は写真に撮って送信してください",
			QuickReply: QuickRepSendAlcPhoto(),
		},
	}
}

func ReplyMoCallHealth() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "疾病、疲労、睡眠不足等はありますか？",
			QuickReply: QuickRepFlagTrue(),
		},
	}
}

func ReplyMCallCarCk() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "車両の日常点検は良好でしたか？",
			QuickReply: QuickRepFlagTrue(),
		},
	}
}

func ReplyMCallNote() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "指示事項、その他必要な事項を入力してください",
			QuickReply: QuickRepMCallNote(),
		},
	}
}

func ReplyECallNote() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "その他必要な事項を入力してください",
			QuickReply: QuickRepECallNote(),
		},
	}
}

func ReplySendMtrPhoto() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "車両のメーターを写真に撮って送信してください",
			QuickReply: QuickRepSendMtrPhoto(),
		},
	}
}

func ReplyDReportTrouble() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "事故、著しい運行の遅延、その他異常な状態が発生した場合は入力してください",
			QuickReply: QuickRepDReportTrouble(),
		},
	}
}

func ReplyDReportNote() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "その他必要な事項を入力してください",
			QuickReply: QuickRepDReportNote(),
		},
	}
}

func ReplyLocNote() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "その他追加する事項を入力してください",
			QuickReply: QuickRepLocNote(),
		},
	}
}

func ReplyResting() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "休憩に入りました\n休憩終了時に位置情報を送信してください",
			QuickReply: QuickRepEndResting(),
		},
	}
}

func ReplyRestNote() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "その他追加する事項を入力してください",
			QuickReply: QuickRepRestNote(),
		},
	}
}

func ReplyWaiting() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "待機に入りました\n待機終了時に位置情報を送信してください",
			QuickReply: QuickRepEndWaiting(),
		},
	}
}

func ReplyWaitNote() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "その他追加する事項を入力してください",
			QuickReply: QuickRepWaitNote(),
		},
	}
}

func ReplyInDuty() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "業務に入りました",
			QuickReply: QuickRepInDuty(),
		},
	}
}

func ReplyDone() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "完了しました",
			QuickReply: QuickRepStartMenu(),
		},
	}
}

func ReplyNone() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "現在進行中の業務はありません",
			QuickReply: QuickRepStartMenu(),
		},
	}
}

func ReplyErrMsg() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "エラーが発生しました\nお手数ですが、スタンプを送信して復帰を試みてください",
		},
	}
}
