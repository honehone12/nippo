package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

func ConfigMenuQuickReply() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{
				Action: UserSettingAction(),
			},
		},
	}
}

func StartMenuQuickReply() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{
				Action: StartDutyAction(),
			},
			{
				Action: UserSettingAction(),
			},
		},
	}
}
