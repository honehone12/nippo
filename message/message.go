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

func validateDst(props *MessageProps, dst string) error {
	org, err := props.db.ReadOrg(props.orgIdStr)
	if err != nil {
		return fmt.Errorf("org is not configured, %w", err)
	}
	if strings.Compare(org.LineId, dst) != 0 {
		return errors.New("dst is not expected, must be misconfigured")
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

	_, err = props.db.ReadUser(sourceId)
	if err != nil {
		log.L.Warnf("user not configured: %s", err)
		return nil
	}

	switch content.Message.(type) {
	default:
		log.L.Warnf("unsupported user message event type: %s\n", content.Type)
		return nil
	}
}

func MessageMain(props *MessageProps, msg *webhook.CallbackRequest) error {
	if err := validateDst(props, msg.Destination); err != nil {
		return err
	}

	for _, e := range msg.Events {
		if err := processEvent(props, e); err != nil {
			log.L.Error(err)
		}
	}

	return nil
}
