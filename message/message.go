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
	OrgIdStr string
	Db       database.Database
}

func validateDst(props *MessageProps, dst string) error {
	org, err := props.Db.ReadOrg(props.OrgIdStr)
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

func processContent(props *MessageProps, content *webhook.MessageEvent) error {
	switch content.Source.(type) {
	case webhook.UserSource:
	default:
		log.L.Warnf("unsupported source: %s\n", content.Source.GetType())
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
