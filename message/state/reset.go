package state

import (
	"nippo/database"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func Reset(props *msgprops.Props, user *schema.User, replyToken string) error {
	err := props.Db().Tx(func(tx database.Transaction) error {
		err := tx.UpdateUserCurrentOpe(user.Id, nil)
		if err != nil {
			return err
		}

		err = tx.UpdateUserCurrentDuty(user.Id, nil)
		if err != nil {
			return err
		}

		return tx.UpdateUserNextState(user.Id, ST_NONE)
	})
	if err != nil {
		return err
	}

	req := messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   reply.ReplyDone(),
	}

	_, err = props.Api().ReplyMessage(&req)
	return err
}
