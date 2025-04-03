package state

import (
	"fmt"
	"nippo/log"
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
	switch e := event.(type) {
	case webhook.MessageEvent:
		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplyDoUserSetting(),
		}

		_, err := props.Api().ReplyMessage(&req)
		return err
	case webhook.PostbackEvent:
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

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyUsInputName(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			return fmt.Errorf("unexpected postback data: %s", e.Postback.Data)
		}
	default:
		log.L.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
