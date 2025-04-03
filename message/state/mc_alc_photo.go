package state

import (
	"nippo/log"
	"nippo/message/blob"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/schema"
	"nippo/sss"

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
	if !ok {
		img, ok := e.Message.(webhook.ImageMessageContent)
		if ok {
			if img.ContentProvider.Type == webhook.ContentProviderTYPE_LINE {
				blobApi := props.BlobApi()
				err := blob.PreRequest(blobApi, img.Id)
				if err != nil {
					return err
				}

				res, err := blobApi.GetMessageContent(img.Id)
				if err != nil {
					return err
				}
				defer res.Body.Close()

				photo, err := sss.Put(props.Ctx(), user.Id, res.Body)
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
					Messages:   reply.ReplyMcHealth(),
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
