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

type StMCallNote struct{}

func (s StMCallNote) Run(
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
			if len(user.Name) <= 1 {
				return errors.New("user does not have name")
			}

			db := props.Db()
			carNum, err := db.ReadMCallCarNum(user.CurrentOperation)
			if err != nil {
				return err
			}
			if len(carNum) == 0 {
				return errors.New("morning call does not have car number")
			}

			text := input.Clean(t.Text)
			err = db.Tx(func(tx database.Transaction) error {
				err := tx.UpdateMCallNote(user.CurrentOperation, text)
				if err != nil {
					return err
				}

				reportId, err := tx.CreateDailyReport(
					user.Id,
					user.Name,
					carNum,
					user.CurrentOperation,
				)
				if err != nil {
					return err
				}

				err = tx.UpdateUserCurrentDuty(user.Id, reportId)
				if err != nil {
					return err
				}

				err = tx.UpdateUserCurrentOpe(user.Id, nil)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_DREPO_M_MTR)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplySendMtr(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyMCallNote(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	default:
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
