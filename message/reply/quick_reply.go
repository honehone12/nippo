package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

func QuickRepConfigMenu() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionUserSetting()},
		},
	}
}

func QuickRepSendAlcPhoto() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionSendAlcPhoto()},
			// {Action: ActionReset()},
		},
	}
}

func QuickRepCarCheck() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionMcCarCheck()},
			// {Action: ActionReset()},
		},
	}
}

func QuickRepInDuty() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionSendGps()},
			{Action: ActionStartResting()},
			{Action: ActionStartWaiting()},
			{Action: ActionStartLoading()},
			{Action: ActionEndDuty()},
		},
	}
}

func QuickRepFlag() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionFlagTrue()},
			// {Action: ActionReset()},
		},
	}
}

func QuickRepStartMenu() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionStartDuty()},
			{Action: ActionUserSetting()},
		},
	}
}
