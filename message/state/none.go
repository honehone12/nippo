package state

import (
	"fmt"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StNone struct{}

func (s StNone) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	switch e := event.(type) {
	case webhook.MessageEvent:
		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplyNone(),
		}

		_, err := props.Api().ReplyMessage(&req)
		return err
	case webhook.PostbackEvent:
		if e.Postback.Data == reply.START_DUTY_ACT_DATA {
			err := props.Db().Once().UpdateUserNextState(user.Id, ST_MCALL_ALC_PHOTO)
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplySendAlcPhoto(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else if e.Postback.Data == reply.USER_SETTING_ACT_DATA {
			err := props.Db().Once().UpdateUserNextState(user.Id, ST_USER_SET_NAME)
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyUserSetName(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			return fmt.Errorf("unexpected postback data: %s", e.Postback.Data)
		}
	default:
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
