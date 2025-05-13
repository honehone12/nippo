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

type StSiteName struct{}

func (s StSiteName) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	e, ok := event.(webhook.MessageEvent)
	if ok {
		t, ok := e.Message.(webhook.TextMessageContent)
		if ok {
			if len(user.CurrentDuty) == 0 {
				return errors.New("user does not have current location")
			}

			text := input.Clean(t.Text)
			err := props.Db().Tx(func(tx database.Transaction) error {
				siteId, err := tx.CreateSite(user.CurrentDuty, text)
				if err != nil {
					return err
				}

				err = tx.UpdateUserCurrentOpe(user.Id, siteId)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_SITE_NOTE)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplySiteNote(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplySiteName(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	} else {
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
