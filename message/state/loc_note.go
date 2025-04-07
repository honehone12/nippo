package state

import (
	"fmt"
	"nippo/log"
	"nippo/message/property"
	"nippo/message/reply"
	"nippo/schema"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type StLocNote struct{}

func (s StLocNote) Run(
	props *property.MsgingProps,
	user *schema.User,
	event webhook.EventInterface,
) error {
	switch e := event.(type) {
	case webhook.PostbackEvent:
		if strings.Compare(e.Postback.Data, reply.IGNORE_TXT_TEMPLATE) == 0 {
			log.L.Warn("ignoring postback data, means shorcut does not work")
			return nil
		} else {
			return fmt.Errorf("unexpected postback data: %s", e.Postback.Data)
		}
	case webhook.MessageEvent:

	default:
		log.L.Warnf("unexpected event: %s", event.GetType())
		return nil
	}
}
