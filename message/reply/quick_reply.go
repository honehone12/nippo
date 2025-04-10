package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

func QuickRepConfigMenu() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionUserSetting()},
		},
	}
}

func QuickRepMCallNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionMCallNote()},
		},
	}
}

func QuickRepECallNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionECallNote()},
		},
	}
}

func QuickRepSendAlcPhoto() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionSendAlcPhoto()},
			{Action: ActionReset()},
		},
	}
}

func QuickRepSendMtrPhoto() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionSendMtrPhoto()},
		},
	}
}

func QuickRepDReportTrouble() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionDReportTrouble()},
		},
	}
}

func QuickRepDReportNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionDReportNote()},
		},
	}
}

func QuickRepLocNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionLocNote()},
		},
	}
}

func QuickRepEndLoading() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionEndLoading()},
		},
	}
}

func QuickRepLoadShipCheck() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionFlagTrue()},
			{Action: ActionFlagFalse()},
		},
	}
}

func QuickRepLoadNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionLoadNote()},
		},
	}
}

func QuickRepEndResting() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionEndResting()},
		},
	}
}

func QuickRepRestNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionRestNote()},
		},
	}
}

func QuickRepEndWaiting() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionEndWaiting()},
		},
	}
}

func QuickRepWaitNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionWaitNote()},
		},
	}
}

func QuickRepInDuty() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionSendLocation()},
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
			{Action: ActionReset()},
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
