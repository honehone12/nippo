package state

import (
	"nippo/database"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func reset(db database.Database, userId []uint8) error {
	return db.Tx(func(tx database.Transaction) error {
		err := tx.UpdateUserCurrentOpe(userId, nil)
		if err != nil {
			return err
		}

		err = tx.UpdateUserCurrentDuty(userId, nil)
		if err != nil {
			return err
		}

		return tx.UpdateUserNextState(userId, ST_NONE)
	})
}

func Decline(props *msgprops.Props, user *schema.User, replyToken string) error {
	err := reset(props.Db(), user.Id)
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

func NoLicense(props *msgprops.Props, user *schema.User, replyToken string) error {
	err := reset(props.Db(), user.Id)
	if err != nil {
		return err
	}

	req := messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   reply.ReplyNoLicense(),
	}

	_, err = props.Api().ReplyMessage(&req)
	return err
}
