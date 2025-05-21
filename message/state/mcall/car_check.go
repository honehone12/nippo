package mcall

import (
	"errors"
	"fmt"
	"nippo/database"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/message/state/index"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type MCallCarCheck struct{}

func (s MCallCarCheck) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	switch e := event.(type) {
	case webhook.MessageEvent:
		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplyMCallCarCheck(),
		}

		_, err := props.Api().ReplyMessage(&req)
		return err
	case webhook.PostbackEvent:
		var checked bool
		switch e.Postback.Data {
		case reply.FLAG_TRUE_ACT_DATA:
			checked = true
		case reply.FLAG_FALSE_ACT_DATA:
			checked = false
		default:
			return fmt.Errorf("unexpected postback data: %s", e.Postback.Data)
		}

		if len(user.CurrentOperation) == 0 {
			return errors.New("user does not have current call")
		}

		err := props.Db().Tx(func(tx database.Transaction) error {
			err := tx.UpdateMCallCarCheck(user.CurrentOperation, checked)
			if err != nil {
				return err
			}

			return tx.UpdateUserNextState(user.Id, index.MCALL_CAR)
		})
		if err != nil {
			return err
		}

		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplyMCallCar(),
		}

		_, err = props.Api().ReplyMessage(&req)
		return err
	default:
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
