package message

import (
	"errors"
	"fmt"
	"nippo/input"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/message/state"
	"nippo/schema"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"gorm.io/gorm"
)

func validateDst(props *msgprops.Props, msg *webhook.CallbackRequest) error {
	db := props.Db()
	orgId := props.OrgId()

	org, err := db.ReadOrg(orgId)
	if err != nil {
		return fmt.Errorf("org is not configured, %w", err)
	}

	cleanDst := input.Clean(msg.Destination)

	if len(org.LineId) == 0 {
		log.L.Warnf("updating org line id to: %s", cleanDst)
		return db.Once().UpdateOrgLineId(orgId, cleanDst)
	}

	if org.LineId != cleanDst {
		return fmt.Errorf(
			"dst has changed from: %s to %s. this is not expected",
			org.LineId, cleanDst,
		)
	}

	return nil
}

func checkBasicContent(event webhook.EventInterface) (string, string, error) {
	var source webhook.SourceInterface
	var repTkn string
	switch e := event.(type) {
	case webhook.MessageEvent:
		source = e.Source
		repTkn = e.ReplyToken
	case webhook.PostbackEvent:
		// add quit now logic here
		// https://github.com/honehone12/nippo/issues/14

		if strings.Compare(e.Postback.Data, reply.IGNORE_TXT_TEMPLATE) == 0 {
			return "", "", nil
		}
		source = e.Source
		repTkn = e.ReplyToken
	default:
		return "", "", fmt.Errorf("unsupported event: %s", event.GetType())
	}

	if source == nil || len(repTkn) == 0 {
		return "", "", errors.New("nil souce or empty reply token")
	}

	s, ok := source.(webhook.UserSource)
	if !ok {
		return "", "", fmt.Errorf("unsupported source: %s", source.GetType())
	}
	if len(s.UserId) == 0 {
		return "", "", errors.New("source has empty user id")
	}

	return s.UserId, repTkn, nil
}

func sendErrMsg(props *msgprops.Props, repTkn string) {
	req := messaging_api.ReplyMessageRequest{
		ReplyToken: repTkn,
		Messages:   reply.ReplyErrMsg(),
	}
	_, err := props.Api().ReplyMessage(&req)
	if err != nil {
		log.L.Errorf("failed to send err msg: %s", err)
	}
}

func processEvent(props *msgprops.Props, event webhook.EventInterface) error {
	source, repTkn, err := checkBasicContent(event)
	if err != nil {
		log.L.Warn(err)
		return nil
	}
	if len(source) == 0 || len(repTkn) == 0 {
		return nil
	}

	cleanSource := input.Clean(source)
	user, err := props.Db().ReadUser(cleanSource)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.L.Warnf("user not configured: %s", err)
		st := state.StUserNotConfigured{}
		unknown := schema.User{LineId: cleanSource}
		err := st.Run(props, &unknown, event)
		if err != nil {
			sendErrMsg(props, repTkn)
			return err
		}

		return nil
	} else if err != nil {
		sendErrMsg(props, repTkn)
		return err
	}

	st, err := state.FromIdx(user.NextState)
	if err != nil {
		sendErrMsg(props, repTkn)
		return err
	}

	err = st.Run(props, user, event)
	if err != nil {
		sendErrMsg(props, repTkn)
		return err
	}

	return nil
}

func MessageMain(props *msgprops.Props, msg *webhook.CallbackRequest) {
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
