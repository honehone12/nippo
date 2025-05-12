package state

import (
	"errors"
	"fmt"
	"nippo/database"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StMCallHealth struct{}

func (s StMCallHealth) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	switch e := event.(type) {
	case webhook.MessageEvent:
		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplyMCallHealth(),
		}

		_, err := props.Api().ReplyMessage(&req)
		return err
	case webhook.PostbackEvent:
		if e.Postback.Data == reply.FLAG_TRUE_ACT_DATA {
			if len(user.CurrentOperation) == 0 {
				return errors.New("user does not have current call")
			}

			err := props.Db().Tx(func(tx database.Transaction) error {
				err := tx.UpdateMCallHealth(user.CurrentOperation)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_MCALL_CAR_CHECK)
			})
			if err != nil {
				return nil
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyMCallCarCheck(),
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
