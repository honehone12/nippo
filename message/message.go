package message

import (
	"errors"
	"fmt"
	"nippo/database"
	"nippo/log"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type MessageProps struct {
	orgIdStr string
	db       database.Database
}

func NewMessageProps(orgIdStr string, db database.Database) *MessageProps {
	return &MessageProps{
		orgIdStr,
		db,
	}
}

func validateDst(props *MessageProps, msg *webhook.CallbackRequest) error {
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

func processEvent(props *MessageProps, event webhook.EventInterface) error {
	switch e := event.(type) {
	case webhook.MessageEvent:
		return processContent(props, &e)
	default:
		log.L.Warnf("unsupported event type: %s\n", event.GetType())
		return nil
	}
}

func getSource(content *webhook.MessageEvent) (string, error) {
	switch c := content.Source.(type) {
	case webhook.UserSource:
		if len(c.UserId) == 0 {

			return "", errors.New("empty source")
		}

		return c.UserId, nil
	default:
		return "", fmt.Errorf("unsupported source: %s", content.Source.GetType())
	}
}

func processContent(props *MessageProps, content *webhook.MessageEvent) error {
	sourceId, err := getSource(content)
	if err != nil {
		log.L.Warn(err)
		return nil
	}

	log.L.Info(sourceId)

	_, err = props.db.ReadUser(sourceId)
	if err != nil {
		log.L.Warnf("user not configured: %s", err)
		return nil
	}

	switch c := content.Message.(type) {
	case webhook.TextMessageContent:
		log.L.Info(c.Text)
		return nil
	default:
		log.L.Warnf("unsupported user message event type: %s\n", content.Type)
		return nil
	}
}

func MessageMain(props *MessageProps, msg *webhook.CallbackRequest) error {
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
