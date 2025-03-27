package message

import (
	"errors"
	"fmt"
	"nippo/database"
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

func MessageMain(props *MessageProps, msg *webhook.CallbackRequest) error {
	if err := validateDst(props, msg.Destination); err != nil {
		return err
	}

	return nil
}
