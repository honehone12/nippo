package reply

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

func QuickRepConfigMenu() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionUsrset()},
		},
	}
}

func QuickRepLicenseExp() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionLicenseExp()},
		},
	}
}

func QuickRepCaller() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionSelf()},
		},
	}
}

func QuickRepMCallNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionMCallNote()},
			{Action: ActionNone()},
		},
	}
}

func QuickRepECallNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionECallNote()},
			{Action: ActionNone()},
		},
	}
}

func QuickRepAlcPhoto() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionAlcPhoto()},
			{Action: ActionDecline()},
		},
	}
}

func QuickRepMtrPhoto() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionMtrPhoto()},
		},
	}
}

func QuickRepDReportTrouble() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionDReportTrouble()},
			{Action: ActionNone()},
		},
	}
}

func QuickRepDReportNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionDReportNote()},
			{Action: ActionNone()},
		},
	}
}

func QuickRepLocNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionLocNote()},
			{Action: ActionNone()},
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
			{Action: ActionNone()},
		},
	}
}

func QuickRepEndResting() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionRestLoc()},
		},
	}
}

func QuickRepSiteNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionSiteNote()},
			{Action: ActionNone()},
		},
	}
}

func QuickRepRestNote() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionRestNote()},
			{Action: ActionNone()},
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
			{Action: ActionNone()},
		},
	}
}

func QuickRepInDuty() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionRegisterSite()},
			{Action: ActionRegisterLocation()},
			{Action: ActionRegisterResting()},
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
			{Action: ActionDecline()},
		},
	}
}

func QuickRepStartMenu() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{Action: ActionStartDuty()},
			{Action: ActionUsrset()},
		},
	}
}
