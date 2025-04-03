package state

import (
	"fmt"
	"nippo/log"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StInDuty struct{}

func (s StInDuty) Run(
	props *property.MessagingProps,
	user *schema.User,
	event webhook.EventInterface,
) error {
	switch e := event.(type) {
	case webhook.MessageEvent:
		_, ok := e.Message.(webhook.LocationMessageContent)
		if ok {
			panic("not implemented")
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyInDuty(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	case webhook.PostbackEvent:
		switch e.Postback.Data {
		case reply.START_WAITING_ACT_DATA:
			panic("not implemented")
		case reply.START_LOADING_ACT_DATA:
			panic("not implemented")
		case reply.START_RESTING_ACT_DATA:
			panic("not implemented")
		case reply.END_DUTY_ACT_DATA:
			err := props.Db().UpdateUserNextState(user.Id, ST_EC_ALC_PHOTO)
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplySendAlcPhoto(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		default:
			return fmt.Errorf("unexpected postback data: %s", e.Postback.Data)
		}
	default:
		log.L.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
