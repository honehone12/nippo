package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

func UserSettingQuickReply() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{
				Action: UserSettingAction(),
			},
		},
	}
}

func StartQuickReply() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{},
	}
}
