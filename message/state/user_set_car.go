package state

import (
	"fmt"
	"nippo/input"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StUserSetCar struct{}

func (s StUserSetCar) Run(
	props *property.MessagingProps,
	user *schema.User,
	event webhook.EventInterface,
) error {
	e, ok := event.(webhook.MessageEvent)
	if !ok {
		return fmt.Errorf("unexpected event: %s", event.GetType())
	}

	t, ok := e.Message.(webhook.TextMessageContent)
	if !ok {
		return fmt.Errorf("unexpected message content: %s", e.Message.GetType())
	}

	db := props.Db()
	text := input.Clean(t.Text)
	err := db.UpdateUserCarNumber(user.Id, text)
	if err != nil {
		return err
	}

	err = db.UpdateUserNextState(user.Id, ST_NONE)
	if err != nil {
		return err
	}

	req := messaging_api.ReplyMessageRequest{
		ReplyToken: e.ReplyToken,
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text:       "完了しました",
				QuickReply: reply.StartQuickReply(),
			},
		},
	}

	_, err = props.Api().ReplyMessage(&req)
	return err
}
