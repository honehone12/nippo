package state

import (
	"fmt"
	"nippo/message/action"
	"nippo/message/property"
	"nippo/schema"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StUserNotConfigured struct {
	idx schema.StateIdx
}

func UserNotConfigured() State {
	return &StUserNotConfigured{idx: ST_USER_NOT_CONFIGURED}
}

func quickReply() *messaging_api.QuickReply {
	return &messaging_api.QuickReply{
		Items: []messaging_api.QuickReplyItem{
			{
				Action: action.UserSettingAction(),
			},
		},
	}
}

func (s *StUserNotConfigured) Run(
	props *property.MessagingProps,
	user *schema.User,
	event webhook.EventInterface,
) error {
	req := messaging_api.ReplyMessageRequest{}
	switch e := event.(type) {
	case webhook.MessageEvent:
		req.ReplyToken = e.ReplyToken
	case webhook.PostbackEvent:
		req.ReplyToken = e.ReplyToken
		if strings.Compare(e.Postback.Data, action.USER_SETTING_ACTION_DATA) == 0 {
			db := props.Db()
			id, err := db.CreateUser(props.OrgId(), user.LineId)
			if err != nil {
				return err
			}

			err = db.UpdateUserNextState(id, ST_USER_SET_NAME)
			if err != nil {
				return err
			}

			req.Messages = []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: "名前を入力してください",
				},
			}

		}

	default:
		return fmt.Errorf("unexpected event: %s", event.GetType())
	}

	req.Messages = []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "ユーザー設定を行ってください",
			QuickReply: quickReply(),
		},
	}

	_, err := props.Api().ReplyMessage(&req)
	return err
}
