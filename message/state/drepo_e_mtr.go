package state

import (
	"errors"
	"nippo/database"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/schema"
	"strconv"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StDRepoEMtr struct{}

func (s StDRepoEMtr) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	e, ok := event.(webhook.MessageEvent)
	if ok {
		text, ok := e.Message.(webhook.TextMessageContent)
		if ok {
			if len(user.CurrentDuty) == 0 {
				return errors.New("user does not have current duty")
			}

			meter, err := strconv.ParseUint(text.Text, 10, 64)
			if err != nil {
				return err
			}

			err = props.Db().Tx(func(tx database.Transaction) error {
				err := tx.UpdateDReportEMtr(user.CurrentDuty, meter)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_DREPO_E_MTR_PHOTO)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplySendMtrPhoto(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplySendMtr(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	} else {
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
