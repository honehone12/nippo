package state

import (
	"errors"
	"nippo/log"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/schema"
	"nippo/sss"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StMcAlcPhoto struct{}

func (s StMcAlcPhoto) Run(
	props *property.MessagingProps,
	user *schema.User,
	event webhook.EventInterface,
) error {
	e, ok := event.(webhook.MessageEvent)
	if ok {
		img, ok := e.Message.(webhook.ImageMessageContent)
		if ok {
			if img.ContentProvider.Type == webhook.ContentProviderTYPE_LINE {
				res, err := props.BlobApi().GetMessageContent(img.Id)
				if err != nil {
					return err
				}
				defer res.Body.Close()

				if !strings.Contains(res.Header.Get("Content-Type"), "image/jpeg") {
					return errors.New("unsupported iamge type")
				}

				photo, err := sss.PutJpg(props.Ctx(), user.Id, res.Body, res.ContentLength)
				if err != nil {
					return err
				}

				db := props.Db()
				callId, err := db.CreateMorningCall(user.Id, photo)
				if err != nil {
					return err
				}

				err = db.UpdateUserCurrentCall(user.Id, callId)
				if err != nil {
					return err
				}

				err = db.UpdateUserNextState(user.Id, ST_MC_HEALTH)
				if err != nil {
					return err
				}

				req := messaging_api.ReplyMessageRequest{
					ReplyToken: e.ReplyToken,
					Messages:   reply.ReplyMoCallHealth(),
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
		log.L.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
