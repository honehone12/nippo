package state

import (
	"errors"
	"nippo/database"
	"nippo/input"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StMCallSetCar struct{}

func (s StMCallSetCar) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	e, ok := event.(webhook.MessageEvent)
	if ok {
		t, ok := e.Message.(webhook.TextMessageContent)
		if ok {
			if len(user.CurrentOperation) == 0 {
				return errors.New("user does not have current call")
			}

			text := input.Clean(t.Text)
			err := props.Db().Tx(func(tx database.Transaction) error {
				err := tx.UpdateMCallCarNumber(user.CurrentOperation, text)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_MCALL_INSPECTION_EXP)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyInspectionExp(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyMCallSetCar(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	} else {
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
