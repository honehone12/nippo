package state

import (
	"errors"
	"fmt"
	"nippo/database"
	"nippo/input"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StECallNote struct{}

func (s StECallNote) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	switch e := event.(type) {
	case webhook.PostbackEvent:
		if e.Postback.Data == reply.IGNORE_TXT_ACT_DATA {
			log.Warn("ignoring postback data, means shorcut does not work")
			return nil
		} else {
			return fmt.Errorf("unexpected postback data: %s", e.Postback.Data)
		}
	case webhook.MessageEvent:
		t, ok := e.Message.(webhook.TextMessageContent)
		if ok {
			if len(user.CurrentOperation) == 0 {
				return errors.New("user does not have current call")
			}

			text := input.Clean(t.Text)
			err := props.Db().Tx(func(tx database.Transaction) error {
				err := tx.UpdateECallNote(user.CurrentOperation, text)
				if err != nil {
					return err
				}

				err = tx.UpdateUserCurrentOpe(user.Id, nil)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_DREPO_E_MTR)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyMtr(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyECallNote(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	default:
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
