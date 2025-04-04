package state

import (
	"errors"
	"fmt"
	"nippo/database"
	"nippo/log"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/schema"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StMcCarCheck struct{}

func (s StMcCarCheck) Run(
	props *property.MessagingProps,
	user *schema.User,
	event webhook.EventInterface,
) error {
	switch e := event.(type) {
	case webhook.MessageEvent:
		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplyMoCallCarCk(),
		}

		_, err := props.Api().ReplyMessage(&req)
		return err
	case webhook.PostbackEvent:
		if strings.Compare(e.Postback.Data, reply.FLAG_TRUE_ACT_DATA) == 0 {
			if len(user.CurrentOperation) == 0 {
				return errors.New("user does not have current call")
			}

			err := props.Db().Tx(func(tx database.Transaction) error {
				err := tx.UpdateMoCallCarCheck(user.CurrentOperation)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_MC_NOTE)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyMoCallNote(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			return fmt.Errorf("unexpected postback data: %s", e.Postback.Data)
		}
	default:
		log.L.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
