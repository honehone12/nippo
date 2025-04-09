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

type StWaiting struct{}

func (s StWaiting) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	e, ok := event.(webhook.MessageEvent)
	if ok {
		l, ok := e.Message.(webhook.LocationMessageContent)
		if ok {
			label := input.Clean(l.Title)
			addr := input.Clean(l.Address)
			err := props.Db().Tx(func(tx database.Transaction) error {
				waitId, err := tx.CreateWaiting(
					user.CurrentDuty,
					label,
					addr,
					l.Latitude,
					l.Longitude,
				)
				if err != nil {
					return err
				}

				err = tx.UpdateUserCurrentOpe(user.Id, waitId)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_WAIT_NOTE)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyWaitNote(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyWaiting(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	} else {
		log.L.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
