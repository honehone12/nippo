package state

import "github.com/line/line-bot-sdk-go/v8/linebot/webhook"

type State interface {
	run(content webhook.MessageContentInterface) (StateIdx, error)
}
