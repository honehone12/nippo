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

type StMCallAlcPhoto struct{}

func (s StMCallAlcPhoto) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	e, ok := event.(webhook.MessageEvent)
	if ok {
		img, ok := e.Message.(webhook.ImageMessageContent)
		if ok {
			if img.ContentProvider.Type == webhook.ContentProviderTYPE_LINE {
				if len(user.Name) <= 1 {
					return errors.New("user does not have name")
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
					callId, err := tx.CreateMorningCall(
						user.Id,
						user.Name,
						photo,
					)
					if err != nil {
						return err
					}

					err = tx.UpdateUserCurrentOpe(user.Id, callId)
					if err != nil {
						return err
					}

					return tx.UpdateUserNextState(user.Id, ST_MCALL_HEALTH)
				})
				if err != nil {
					return err
				}

				req := messaging_api.ReplyMessageRequest{
					ReplyToken: e.ReplyToken,
					Messages:   reply.ReplyMCallHealth(),
				}

				_, err = props.Api().ReplyMessage(&req)
				return err
			}
		}

		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplyAlcPhoto(),
		}

		_, err := props.Api().ReplyMessage(&req)
		return err
	} else {
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
