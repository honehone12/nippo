package exit

import (
	"nippo/database"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/message/state/index"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func exit(db database.Database, userId []uint8) error {
	return db.Tx(func(tx database.Transaction) error {
		err := tx.UpdateUserCurrentOpe(userId, nil)
		if err != nil {
			return err
		}

		err = tx.UpdateUserCurrentDuty(userId, nil)
		if err != nil {
			return err
		}

		return tx.UpdateUserNextState(userId, index.NONE)
	})
}

func Decline(props *msgprops.Props, user *schema.User, replyToken string) error {
	err := exit(props.Db(), user.Id)
	if err != nil {
		return err
	}

	req := messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   reply.ReplyDecline(),
	}

	_, err = props.Api().ReplyMessage(&req)
	return err
}

func Expired(props *msgprops.Props, user *schema.User, replyToken string) error {
	err := exit(props.Db(), user.Id)
	if err != nil {
		return err
	}

	req := messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   reply.ReplyExpired(),
	}

	_, err = props.Api().ReplyMessage(&req)
	return err
}
