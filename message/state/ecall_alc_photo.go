package state

import (
	"errors"
	"nippo/database"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/schema"
	"nippo/sss"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StECallAlcPhoto struct{}

func (s StECallAlcPhoto) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	e, ok := event.(webhook.MessageEvent)
	if ok {
		img, ok := e.Message.(webhook.ImageMessageContent)
		if ok {
			if img.ContentProvider.Type == webhook.ContentProviderTYPE_LINE {
				if len(user.CurrentDuty) == 0 {
					return errors.New("user does not have current duty")
				}
				if len(user.Name) <= 1 {
					return errors.New("user does not have name")
				}
				if len(user.CarNumber) == 0 {
					return errors.New("user does not have car number")
				}

				res, err := props.BlobApi().GetMessageContent(img.Id)
				if err != nil {
					return err
				}
				defer res.Body.Close()

				if res.Header.Get("Content-Type") != "image/jpeg" {
					return errors.New("unsupported iamge type")
				}

				photo, err := sss.PutJpg(props, user.Id, res.Body, res.ContentLength)
				if err != nil {
					return err
				}

				err = props.Db().Tx(func(tx database.Transaction) error {
					callId, err := tx.CreateEveningCall(
						user.Id,
						user.Name,
						user.CarNumber,
						photo,
					)
					if err != nil {
						return err
					}

					err = tx.UpdateUserCurrentOpe(user.Id, callId)
					if err != nil {
						return err
					}

					err = tx.UpdateDReportECall(user.CurrentDuty, callId)
					if err != nil {
						return err
					}

					return tx.UpdateUserNextState(user.Id, ST_ECALL_NOTE)
				})
				if err != nil {
					return err
				}

				req := messaging_api.ReplyMessageRequest{
					ReplyToken: e.ReplyToken,
					Messages:   reply.ReplyECallNote(),
				}

				_, err = props.Api().ReplyMessage(&req)
				return err
			}
		}

		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplySendAlcPhoto(),
		}

		_, err := props.Api().ReplyMessage(&req)
		return err
	} else {
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
