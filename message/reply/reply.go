package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

func ReplyDoUsrset() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "ユーザー設定を行ってください",
			QuickReply: QuickRepConfigMenu(),
		},
	}
}

func ReplyUsrsetName() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "名前を入力してください",
		},
	}
}

func ReplyUsrsetEmail() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "メールアドレスを『英数字』で入力してください",
		},
	}
}

func ReplyLicenseExp() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "『免許証』の有効期限を入力してください",
			QuickReply: QuickRepLicenseExp(),
		},
	}
}

func ReplyInspectionExp() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "『車検証』の有効期限を入力してください",
			QuickReply: QuickRepLicenseExp(),
		},
	}
}

func ReplyAlcPhoto() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "アルコール検査機を使用し、結果が良好な場合は写真に撮って送信してください",
			QuickReply: QuickRepAlcPhoto(),
		},
	}
}

func ReplyCaller() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "点呼者を入力してください",
			QuickReply: QuickRepCaller(),
		},
	}
}

func ReplyMCallHealth() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "疾病、疲労、睡眠不足等はありますか？",
			QuickReply: QuickRepFlag(),
		},
	}
}

func ReplyMCallCarCheck() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "車両の日常点検は良好でしたか？",
			QuickReply: QuickRepFlag(),
		},
	}
}

func ReplyMCallCar() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "車両番号を入力してください",
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

func ReplyMtr() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "車両のメーターを『英数字』、『自然数』で入力してください",
		},
	}
}

func ReplyMtrPhoto() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "車両のメーターを写真に撮って送信してください",
			QuickReply: QuickRepMtrPhoto(),
		},
	}
}

func ReplyDReportNonDutyDst() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "空走行距離を『英数字』、『自然数』で入力してください",
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

func ReplySiteName() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: "現場名を入力してください",
		},
	}
}

func ReplySiteNote() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "その他追加する事項を入力してください",
			QuickReply: QuickRepSiteNote(),
		},
	}
}

func ReplyLoading() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "荷役（附帯）作業に入りました\n『終了時』に位置情報を送信してください",
			QuickReply: QuickRepEndLoading(),
		},
	}
}

func ReplyLoadShipCheck() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "作業に関して荷主の確認は得られましたか？",
			QuickReply: QuickRepLoadShipCheck(),
		},
	}
}

func ReplyLoadNote() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "その他追加する事項を入力してください",
			QuickReply: QuickRepLoadNote(),
		},
	}
}

func ReplyRestLoc() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "位置情報を送信してください",
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
			Text:       "待機に入りました\n『終了時』に位置情報を送信してください",
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

func ReplyDecline() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "業務を辞退しました\n必ず管理者に連絡してください",
			QuickReply: QuickRepStartMenu(),
		},
	}
}

func ReplyExpired() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "有効期限が切れています\n必ず管理者に連絡してください",
			QuickReply: QuickRepStartMenu(),
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

func ReplyGoodJob() []messaging_api.MessageInterface {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "完了しました\nおつかれさまです！(^o^)丿",
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
			Text: "エラーが発生しました( ﾟДﾟ)\nお手数ですが、スタンプを送信して復帰を試みてください",
		},
	}
}
