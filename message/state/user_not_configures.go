package state

import (
	"fmt"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/schema"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StUserNotConfigured struct{}

func (s StUserNotConfigured) Run(
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

		if strings.Compare(e.Postback.Data, reply.USER_SETTING_ACT_DATA) == 0 {
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

			_, err = props.Api().ReplyMessage(&req)
			return err
		}
	default:
		return fmt.Errorf("unexpected event: %s", event.GetType())
	}

	req.Messages = []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text:       "ユーザー設定を行ってください",
			QuickReply: reply.UserSettingQuickReply(),
		},
	}

	_, err := props.Api().ReplyMessage(&req)
	return err
}
