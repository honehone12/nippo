package state

import (
	"errors"
	"nippo/message/property"
	"nippo/schema"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type State interface {
	Run(
		props *property.MsgProps,
		event webhook.EventInterface,
	) error
}

func FromIdx(idx schema.StateIdx) (State, error) {
	switch idx {
	case schema.ST_IDX_NONE:
		return &StNone{idx}, nil
	case schema.ST_USER_NOT_CONFIGURED:
		return &StUserNotConfigured{idx}, nil
	default:
		return nil, errors.New("undefined state index")
	}
}

// func __processEvent(props *property.MessagingProperties, event *webhook.MessageEvent) error {
// 	sourceId, err := getSource(event.Source)
// 	if err != nil {
// 		log.L.Warn(err)
// 		return nil
// 	}

// 	log.L.Info(sourceId)

// 	switch c := event.Message.(type) {
// 	case webhook.TextMessageContent:
// 		log.L.Info(c.Text)
// 		return nil
// 	default:
// 		log.L.Warnf("unsupported content type: %s\n", event.Type)
// 		return nil
// 	}
// }
