package state

import (
	"errors"
	"fmt"
	"nippo/database"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/schema"
	"nippo/timeconv"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StMCallLicenseExp struct{}

func (s StMCallLicenseExp) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	switch e := event.(type) {
	case webhook.MessageEvent:
		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplyLicenseExp(),
		}

		_, err := props.Api().ReplyMessage(&req)
		return err
	case webhook.PostbackEvent:
		if e.Postback.Data == reply.EXP_DATE_ACT_DATA {
			if len(user.CurrentOperation) == 0 {
				return errors.New("user does not have current call")
			}

			d, ok := e.Postback.Params["date"]
			if !ok {
				return errors.New("unexpected postback without date")
			}

			exp, err := timeconv.ParseDateInJst(d)
			if err != nil {
				return err
			}
			if timeconv.NowJst().After(exp) {
				return Expired(props, user, e.ReplyToken)
			}

			err = props.Db().Tx(func(tx database.Transaction) error {
				err := tx.UpdateMCallLicenseExp(user.CurrentOperation, exp.UTC())
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_MCALL_CALLER)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyCaller(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			return fmt.Errorf("unexpected postback data: %s", e.Postback.Data)
		}
	default:
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
