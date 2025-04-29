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

type StLoading struct{}

func (s StLoading) Run(
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
				loadId, err := tx.CreateLoading(
					user.CurrentDuty,
					label,
					addr,
					l.Latitude,
					l.Longitude,
				)
				if err != nil {
					return err
				}

				err = tx.UpdateUserCurrentOpe(user.Id, loadId)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_LOAD_SHIP_CHECK)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyLoadShipCheck(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyLoading(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	} else {
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
