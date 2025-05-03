package admin

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"nippo/admin/adminprops"
	"nippo/cognito"
	"nippo/schema"

	"github.com/google/uuid"
)

func Invite(props *adminprops.Props, form url.Values) ([]byte, error) {
	userIdStr := form.Get("user")
	if len(userIdStr) == 0 {
		return nil, errors.New("invalid form")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return nil, err
	}

	invitedUser, err := props.Db().ReadUserEmailOrgId(userId[:])
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(invitedUser.OrgId, props.OrgId()) {
		return nil, errors.New("wrong org")
	}
	if len(invitedUser.Email) == 0 {
		return nil, errors.New("user email is empty")
	}

	err = cognito.Invite(props, invitedUser.Email)
	if err != nil {
		return nil, err
	}

	invited := schema.Invited{
		UserEmail: invitedUser.Email,
	}

	return json.Marshal(invited)
}
