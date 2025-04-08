package state

import (
	"errors"
	"fmt"
	"nippo/database"
	"nippo/input"
	"nippo/log"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StInDuty struct{}

func (s StInDuty) Run(
	props *property.MsgingProps,
	user *schema.User,
	event webhook.EventInterface,
) error {
	switch e := event.(type) {
	case webhook.MessageEvent:
		l, ok := e.Message.(webhook.LocationMessageContent)
		if ok {
			if len(user.CurrentDuty) == 0 {
				return errors.New("user does not have current duty")
			}

			label := input.Clean(l.Title)
			addr := input.Clean(l.Address)
			err := props.Db().Tx(func(tx database.Transaction) error {
				locId, err := tx.CreateLocation(
					user.CurrentDuty,
					label,
					addr,
					l.Latitude,
					l.Longitude,
				)
				if err != nil {
					return err
				}

				err = tx.UpdateUserCurrentOpe(user.Id, locId)
				if err != nil {
					return err
				}

				return tx.UpdateUserNextState(user.Id, ST_LOC_NOTE)
			})
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyLocNote(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		} else {
			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyInDuty(),
			}

			_, err := props.Api().ReplyMessage(&req)
			return err
		}
	case webhook.PostbackEvent:
		switch e.Postback.Data {
		case reply.START_WAITING_ACT_DATA:
			err := props.Db().Once().UpdateUserNextState(user.Id, ST_WAITING)
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyWaiting(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		case reply.START_LOADING_ACT_DATA:
			panic("not implemented")
		case reply.START_RESTING_ACT_DATA:
			err := props.Db().Once().UpdateUserNextState(user.Id, ST_RESTING)
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplyResting(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		case reply.END_DUTY_ACT_DATA:
			err := props.Db().Once().UpdateUserNextState(user.Id, ST_ECALL_ALC_PHOTO)
			if err != nil {
				return err
			}

			req := messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   reply.ReplySendAlcPhoto(),
			}

			_, err = props.Api().ReplyMessage(&req)
			return err
		default:
			return fmt.Errorf("unexpected postback data: %s", e.Postback.Data)
		}
	default:
		log.L.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
