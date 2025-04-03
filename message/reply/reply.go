package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

func ReplyDoUserSetting() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "ユーザー設定を行ってください",
			QuickReply: QuickReplyConfigMenu(),
		},
	}
}

func ReplyUsInputName() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "名前を入力してください",
		},
	}
}

func ReplyUsInputCar() []messaging_api.MessageInterface {
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
			QuickReply: QuickReplySendAlcPhoto(),
		},
	}
}

func ReplyMcHealth() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "疾病、疲労、睡眠不足等はありますか？",
			QuickReply: QuickReplyFlag(),
		},
	}
}

func ReplyMcCarCheck() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "車両の日常点検は良好でしたか？",
			QuickReply: QuickReplyCarCheck(),
		},
	}
}

func ReplyMcNote() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "指示事項、その他必要な事項を入力してください",
		},
	}
}

func ReplyInDuty() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "業務中に入りました",
			// QuickReply: ,
		},
	}
}

func ReplyDone() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "完了しました",
			QuickReply: QuickReplyStartMenu(),
		},
	}
}

func ReplyNone() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "現在進行中のタスクはありません",
			QuickReply: QuickReplyStartMenu(),
		},
	}
}

func ReplyErrMsg() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "エラーが発生しております\nお手数ですが、スタンプを送信して復帰を試みてください",
		},
	}
}
