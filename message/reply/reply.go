package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

func ReplyDoUserSetting() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "ユーザー設定を行ってください",
			QuickReply: ConfigMenuQuickReply(),
		},
	}
}

func ReplyInputName() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "名前を入力してください",
		},
	}
}

func ReplyInputCar() []messaging_api.MessageInterface {
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
			QuickReply: SendAlcPhotoQuickReply(),
		},
	}
}

func ReplyDone() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "完了しました",
			QuickReply: StartMenuQuickReply(),
		},
	}
}

func ReplyNone() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "現在進行中のタスクはありません",
			QuickReply: StartMenuQuickReply(),
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
