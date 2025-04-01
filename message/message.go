package message

import (
	"errors"
	"fmt"
	"nippo/log"
	"nippo/message/property"
	"nippo/message/state"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"gorm.io/gorm"
)

func validateDst(props *property.MessagingProperties, msg *webhook.CallbackRequest) error {
	org, err := props.Db().ReadOrg(props.OrgId())
	if err != nil {
		return fmt.Errorf("org is not configured, %w", err)
	}

	if len(org.LineId) == 0 {
		log.L.Warnf("updating org line id to: %s\n", msg.Destination)
		return nil
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

func processEvent(props *property.MessagingProperties, event webhook.EventInterface) error {
	source, err := getSource(event)
	if err != nil {
		return err
	}

	user, err := props.Db().ReadUser(source)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.L.Warnf("user not configured: %s", err)

		return nil
	} else if err != nil {
		return err
	}

	st, err := state.FromIdx(user.NextState)
	if err != nil {
		return err
	}

	return st.Run(props, event)
}

func MessageMain(props *property.MessagingProperties, msg *webhook.CallbackRequest) error {
	if err := validateDst(props, msg); err != nil {
		return err
	}

	for _, e := range msg.Events {
		if err := processEvent(props, e); err != nil {
			log.L.Error(err)
		}
	}

	return nil
}
