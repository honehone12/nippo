package state

import (
	"errors"
	"nippo/database"
	"nippo/log"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/schema"
	"nippo/sss"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StDRepoMtrPhoto struct{}

func (s StDRepoMtrPhoto) Run(
	props *property.MsgingProps,
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

				res, err := props.BlobApi().GetMessageContent(img.Id)
				if err != nil {
					return err
				}
				defer res.Body.Close()

				if !strings.Contains(res.Header.Get("Content-Type"), "image/jpeg") {
					return errors.New("unsupported iamge type")
				}

				photo, err := sss.PutJpg(props, user.Id, res.Body, res.ContentLength)
				if err != nil {
					return err
				}

				err = props.Db().Tx(func(tx database.Transaction) error {
					err := tx.UpdateDReportMtrPhoto(user.CurrentDuty, photo)
					if err != nil {
						return err
					}

					return tx.UpdateUserNextState(user.Id, ST_DREPO_TROUBLE)
				})
				if err != nil {
					return err
				}

				req := messaging_api.ReplyMessageRequest{
					ReplyToken: e.ReplyToken,
					Messages:   reply.ReplyDReportTrouble(),
				}

				_, err = props.Api().ReplyMessage(&req)
				return err
			}
		}

		req := messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages:   reply.ReplySendMtrPhoto(),
		}

		_, err := props.Api().ReplyMessage(&req)
		return err
	} else {
		log.L.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
