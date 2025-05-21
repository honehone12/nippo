package loading

import (
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

type Loading struct{}

func (s Loading) Run(
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

				return tx.UpdateUserNextState(user.Id, index.LOAD_SHIP_CHECK)
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
