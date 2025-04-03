package message

import (
	"errors"
	"fmt"
	"nippo/log"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/message/state"
	"nippo/schema"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"gorm.io/gorm"
)

func validateDst(props *property.MessagingProps, msg *webhook.CallbackRequest) error {
	db := props.Db()
	orgId := props.OrgId()

	org, err := db.ReadOrg(orgId)
	if err != nil {
		return fmt.Errorf("org is not configured, %w", err)
	}

	if len(org.LineId) == 0 {
		log.L.Warnf("updating org line id to: %s", msg.Destination)
		return db.UpdateOrgLineId(orgId, msg.Destination)
	}

	if strings.Compare(org.LineId, msg.Destination) != 0 {
		return fmt.Errorf(
			"dst has changed from: %s to %s. this is not expected",
			org.LineId, msg.Destination,
		)
	}

	return nil
}

func getSource(event webhook.EventInterface) (string, error) {
	var source webhook.SourceInterface
	switch e := event.(type) {
	case webhook.MessageEvent:
		source = e.Source
	case webhook.PostbackEvent:
		source = e.Source
	default:
		return "", fmt.Errorf("unsupported event: %s", event.GetType())
	}
	if source == nil {
		return "", errors.New("nil source")
	}

	s, ok := source.(webhook.UserSource)
	if !ok {
		return "", fmt.Errorf("unsupported source: %s", source.GetType())
	}
	if len(s.UserId) == 0 {
		return "", errors.New("empty source")
	}

	return s.UserId, nil
}

func sendErrMsg(props *property.MessagingProps, event webhook.EventInterface) {
	req := messaging_api.ReplyMessageRequest{}
	switch e := event.(type) {
	case webhook.MessageEvent:
		req.ReplyToken = e.ReplyToken
	case webhook.PostbackEvent:
		req.ReplyToken = e.ReplyToken
	default:
		log.L.Warn("stop sending err msg as unexpected msg type")
		return
	}

	req.Messages = reply.ReplyErrMsg()
	_, err := props.Api().ReplyMessage(&req)
	if err != nil {
		log.L.Errorf("failed to send err msg: %s", err)
	}
}

func processEvent(props *property.MessagingProps, event webhook.EventInterface) error {
	source, err := getSource(event)
	if err != nil {
		log.L.Warn(err)
		return nil
	}

	user, err := props.Db().ReadUser(source)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.L.Warnf("user not configured: %s", err)
		st := state.StUserNotConfigured{}

		err = st.Run(props, &schema.User{LineId: source}, event)
	}
	if err != nil {
		sendErrMsg(props, event)
		return err
	}

	// add quit now logic here
	// https://github.com/honehone12/nippo/issues/14

	st, err := state.FromIdx(user.NextState)
	if err != nil {
		sendErrMsg(props, event)
		return err
	}

	err = st.Run(props, user, event)
	if err != nil {
		sendErrMsg(props, event)
		return err
	}

	return nil
}

func MessageMain(props *property.MessagingProps, msg *webhook.CallbackRequest) {
	if err := validateDst(props, msg); err != nil {
		log.L.Error(err)
		return
	}

	for _, e := range msg.Events {
		if err := processEvent(props, e); err != nil {
			log.L.Error(err)
		}
	}
}
