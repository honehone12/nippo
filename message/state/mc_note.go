package state

import (
	"nippo/input"
	"nippo/log"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StMcNote struct{}

func (s StMcNote) Run(
	props *property.MessagingProps,
	user *schema.User,
	event webhook.EventInterface,
) error {
	e, ok := event.(webhook.MessageEvent)
	if ok {
		t, ok := e.Message.(webhook.TextMessageContent)
		if ok {
			db := props.Db()
			text := input.Clean(t.Text)

			err := db.UpdateMoCallNote(user.Id, text)
			if err != nil {
				return err
			}

			err = db.UpdateUserCurrentCall(user.Id, nil)
			if err != nil {
				return err
			}

			err = db.UpdateUserNextState(user.Id, ST_IN_DUTY)
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyInDuty(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyMcNote(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	} else {
		log.L.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
