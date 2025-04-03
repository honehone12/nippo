package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

func QuickReplyConfigMenu() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionUserSetting()},
		},
	}
}

func QuickReplySendAlcPhoto() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionSendAlcPhoto()},
			// {Action: ActionReset()},
		},
	}
}

func QuickReplyCarCheck() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionFlagTrueAndMcNote()},
			// {Action: ActionReset()},
		},
	}
}

func QuickReplyFlag() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionFlagTrue()},
			// {Action: ActionReset()},
		},
	}
}

func QuickReplyStartMenu() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionStartDuty()},
			{Action: ActionUserSetting()},
		},
	}
}
