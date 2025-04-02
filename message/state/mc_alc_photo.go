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
				err := blob.PreRequest(props.BlobApi(), img.Id)
				if err != nil {
					return err
				}

				res, err := props.BlobApi().GetMessageContent(img.Id)
				if err != nil {
					return err
				}
				defer res.Body.Close()

				err = sss.Put(props.Ctx(), user.Id, res.Body)
				if err != nil {
					return err
				}

				err = props.Db().UpdateUserNextState(user.Id, ST_MC_HEALTH)
				if err != nil {
					return err
				}

			}
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplySendAlcPhoto(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	} else {
		log.L.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
