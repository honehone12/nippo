package resting

import (
	"errors"
	"nippo/database"
	"nippo/input"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/message/state/index"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type Resting struct{}

func (s Resting) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	e, ok := event.(webhook.MessageEvent)
	if ok {
		l, ok := e.Message.(webhook.LocationMessageContent)
		if ok {
			if len(user.CurrentDuty) == 0 {
				return errors.New("user does not have current duty")
			}

			label := input.Clean(l.Title)
			addr := input.Clean(l.Address)
			err := props.Db().Tx(func(tx database.Transaction) error {
				restId, err := tx.CreateResting(
					user.CurrentDuty,
					label,
					addr,
					l.Latitude,
					l.Longitude,
				)
				if err != nil {
					return err
				}

				err = tx.UpdateUserCurrentOpe(user.Id, restId)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, index.REST_NOTE)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyRestNote(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyRestLoc(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	} else {
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
