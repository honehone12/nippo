package blob

import (
	"errors"
	"nippo/log"
	"time"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func PreRequest(blobApi *messaging_api.MessagingApiBlobAPI, blobId string) error {
	// default lambda time out is 3s
	// try within 2s / 250 = 8
	const WAIT_MIL = 250
	const MAX_TRY = 8
	try := 0
LOOP:
	for {
		if try >= MAX_TRY {
			return errors.New("exceeded max tries, failed to obtain img")
		}

		res, err := blobApi.GetMessageContentTranscodingByMessageId(blobId)
		if err != nil {
			return err
		}

		switch res.Status {
		case messaging_api.GetMessageContentTranscodingResponseSTATUS_SUCCEEDED:
			break LOOP
		case messaging_api.GetMessageContentTranscodingResponseSTATUS_PROCESSING:
			log.L.Warn("being going to sleep")
			time.Sleep(time.Millisecond * WAIT_MIL)
			try++
		default:
			return errors.New("img processing failed")
		}
	}

	return nil
}
