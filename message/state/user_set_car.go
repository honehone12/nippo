package state

import (
	"nippo/database"
	"nippo/input"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StUserSetCar struct{}

func (s StUserSetCar) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	e, ok := event.(webhook.MessageEvent)
	if ok {
		t, ok := e.Message.(webhook.TextMessageContent)
		if ok {
			text := input.Clean(t.Text)
			err := props.Db().Tx(func(tx database.Transaction) error {
				err := tx.UpdateUserCarNumber(user.Id, text)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_NONE)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyDone(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyUserSetCar(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	} else {
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
