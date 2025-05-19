package state

import (
	"errors"
	"fmt"
	"nippo/database"
	"nippo/input"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StECallCaller struct{}

func (s StECallCaller) Run(
	props *msgprops.Props,
	user *schema.User,
	event webhook.EventInterface,
) error {
	var caller schema.CleanText
	var repTkn string

	switch e := event.(type) {
	case webhook.MessageEvent:
		t, ok := e.Message.(webhook.TextMessageContent)
		if ok {
			text := input.Clean(t.Text)
			caller = text
			repTkn = e.ReplyToken
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyCaller(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	case webhook.PostbackEvent:
		if e.Postback.Data == reply.SELF_ACT_DATA {
			caller = user.Name
			repTkn = e.ReplyToken
		} else {
			return fmt.Errorf("unexpected postback data: %s", e.Postback.Data)
		}
	default:
		log.Warnf("unexpected event: %s", event.GetType())
		return nil
	}

	if len(user.CurrentOperation) == 0 {
		return errors.New("user does not have current operation")
	}

	err := props.Db().Tx(func(tx database.Transaction) error {
		err := tx.UpdateECallCaller(user.CurrentOperation, caller)
		if err != nil {
			return err
		}

		return tx.UpdateUserNextState(user.Id, ST_ECALL_NOTE)
	})
	if err != nil {
		return err
	}

	req := messaging_api.ReplyMessageRequest{
		ReplyToken: repTkn,
		Messages:   reply.ReplyMCallNote(),
	}

	_, err = props.Api().ReplyMessage(&req)
	return err
}
