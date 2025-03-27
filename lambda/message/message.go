package message

import (
	"nippo/log"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func MessageMain(msg *webhook.CallbackRequest) error {

	log.L.Infof("%#v\n", msg)

	return nil
}
