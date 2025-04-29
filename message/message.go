package message

import (
	"bytes"
	"errors"
	"fmt"
	"nippo/input"
	"nippo/log"
	"nippo/message/msgprops"
	"nippo/message/reply"
	"nippo/message/state"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"gorm.io/gorm"
)

type basicContent struct {
	source     string
	replyToken string
	reset      bool
}

func (bc basicContent) ignore() bool {
	return len(bc.source) == 0 || len(bc.replyToken) == 0
}

func validateDst(props *msgprops.Props, msg *webhook.CallbackRequest) error {
	db := props.Db()
	orgId := props.OrgId()

	org, err := db.ReadOrg(orgId)
	if err != nil {
		return fmt.Errorf("org is not configured, %w", err)
	}

	cleanDst := input.Clean(msg.Destination)

	if len(org.LineId) == 0 {
		log.Warnf("updating org line id to: %s", cleanDst)
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

func checkBasicContent(event webhook.EventInterface) (*basicContent, error) {
	var source webhook.SourceInterface
	var content basicContent
	switch e := event.(type) {
	case webhook.MessageEvent:
		source = e.Source
		content.replyToken = e.ReplyToken
	case webhook.PostbackEvent:
		switch e.Postback.Data {
		case reply.IGNORE_TXT_ACT_DATA:
			return &content, nil
		case reply.RESET_ACT_DATA:
			source = e.Source
			content.replyToken = e.ReplyToken
			content.reset = true
		default:
			source = e.Source
			content.replyToken = e.ReplyToken
		}
	default:
		return nil, fmt.Errorf("unsupported event: %s", event.GetType())
	}

	if len(content.replyToken) == 0 {
		return nil, errors.New("empty reply token")
	}

	s, ok := source.(webhook.UserSource)
	if !ok {
		return nil, fmt.Errorf("unsupported source: %s", source.GetType())
	}
	if len(s.UserId) == 0 {
		return nil, errors.New("source has empty user id")
	}

	content.source = s.UserId
	return &content, nil
}

func sendErrMsg(props *msgprops.Props, repTkn string) {
	req := messaging_api.ReplyMessageRequest{
		ReplyToken: repTkn,
		Messages:   reply.ReplyErrMsg(),
	}
	_, err := props.Api().ReplyMessage(&req)
	if err != nil {
		log.Errorf("failed to send err msg: %s", err)
	}
}

func processEvent(props *msgprops.Props, event webhook.EventInterface) (*basicContent, error) {
	bc, err := checkBasicContent(event)
	if err != nil {
		log.Warn(err)
		return nil, nil
	}
	if bc.ignore() {
		return nil, nil
	}

	cleanSource := input.Clean(bc.source)
	user, err := props.Db().ReadUser(cleanSource)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Warnf("user not configured: %s", err)
		st := state.StUserNotConfigured{}
		unknown := schema.User{LineId: cleanSource}
		err := st.Run(props, &unknown, event)
		if err != nil {
			return bc, err
		}

		return nil, nil
	} else if err != nil {
		return bc, err
	}

	if !bytes.Equal(props.OrgId(), user.OrgId) {
		return bc, errors.New("org id do not match")
	}

	if bc.reset {
		err := state.Reset(props, user, bc.replyToken)
		if err != nil {
			return bc, err
		}
		return nil, nil
	}

	st, err := state.FromIdx(user.NextState)
	if err != nil {
		return bc, err
	}

	err = st.Run(props, user, event)
	if err != nil {
		return bc, err
	}

	return nil, nil
}

func eventMain(props *msgprops.Props, event webhook.EventInterface) error {
	if bc, err := processEvent(props, event); err != nil {
		sendErrMsg(props, bc.replyToken)
		return err
	}

	return nil
}

func MessageMain(props *msgprops.Props, msg *webhook.CallbackRequest) {
	if err := validateDst(props, msg); err != nil {
		log.Error(err)
		return
	}

	for _, e := range msg.Events {
		if err := eventMain(props, e); err != nil {
			log.Error(err)
		}
	}
}
