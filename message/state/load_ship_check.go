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

type StLoadShipCheck struct{}

func (s StLoadShipCheck) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	switch e := event.(type) {
	case webhook.MessageEvent:
		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplyLoadShipCheck(),
		}

		_, err := props.Api().ReplyMessage(&req)
		return err
	case webhook.PostbackEvent:
		if len(user.CurrentOperation) == 0 {
			return errors.New("user does not have current loading")
		}

		var check bool
		switch e.Postback.Data {
		case reply.FLAG_TRUE_ACT_DATA:
			check = true
		case reply.FLAG_FALSE_ACT_DATA:
			check = false
		default:
			return fmt.Errorf("unexpected postback data: %s", e.Postback.Data)
		}

		err := props.Db().Tx(func(tx database.Transaction) error {
			err := tx.UpdateLoadShipCheck(user.CurrentOperation, check)
			if err != nil {
				return nil
			}

			return tx.UpdateUserNextState(user.Id, ST_LOAD_NOTE)
		})
		if err != nil {
			return err
		}

		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplyLoadNote(),
		}

		_, err = props.Api().ReplyMessage(&req)
		return err
	default:
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
