package message

import (
	"errors"
	"fmt"
	"nippo/log"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"gorm.io/gorm"
)

func validateDst(props *MessagingProperties, msg *webhook.CallbackRequest) error {
	org, err := props.db.ReadOrg(props.orgIdStr)
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

func getSource(source webhook.SourceInterface) (string, error) {
	s, ok := source.(webhook.UserSource)
	if !ok {
		return "", fmt.Errorf("unsupported source: %s", source.GetType())
	}
	if len(s.UserId) == 0 {
		return "", errors.New("empty source")
	}

	return s.UserId, nil
}

func processMessageEvent(props *MessagingProperties, event *webhook.MessageEvent) error {
	sourceId, err := getSource(event.Source)
	if err != nil {
		log.L.Warn(err)
		return nil
	}

	log.L.Info(sourceId)

	_, err = props.db.ReadUser(sourceId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.L.Warnf("user not configured: %s", err)

		return nil
	} else if err != nil {
		return err
	}

	switch c := event.Message.(type) {
	case webhook.TextMessageContent:
		log.L.Info(c.Text)
		return nil
	default:
		log.L.Warnf("unsupported content type: %s\n", event.Type)
		return nil
	}
}

func processPostbackEvent(props *MessagingProperties, event *webhook.PostbackEvent) error {
	sourceId, err := getSource(event)
	if err != nil {
		log.L.Warn(err)
		return nil
	}

	log.L.Info(sourceId)

	panic("todo")
}

func processEvent(
	props *MessagingProperties,
	event webhook.EventInterface,
) error {
	switch e := event.(type) {
	case webhook.MessageEvent:
		return processMessageEvent(props, &e)
	case webhook.PostbackEvent:
		panic("todo")
	default:
		log.L.Warnf("unsupported event type: %s\n", event.GetType())
		return nil
	}
}

func MessageMain(props *MessagingProperties, msg *webhook.CallbackRequest) error {
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
