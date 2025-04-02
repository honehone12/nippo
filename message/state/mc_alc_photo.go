package state

import (
	"errors"
	"nippo/log"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/schema"
	"nippo/sss"
	"time"

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
				try := 1
			LOOP:
				for {
					if try > 8 {
						return errors.New("exceeded max tries, failed to obtain img")
					}

					res, err := blobApi.GetMessageContentTranscodingByMessageId(img.Id)
					if err != nil {
						return err
					}

					switch res.Status {
					case messaging_api.GetMessageContentTranscodingResponseSTATUS_SUCCEEDED:
						break LOOP
					case messaging_api.GetMessageContentTranscodingResponseSTATUS_PROCESSING:
						log.L.Warn("being going to sleep")
						time.Sleep(time.Millisecond * 250)
						try++
					default:
						return errors.New("img processing failed")
					}
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
