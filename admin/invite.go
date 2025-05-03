package admin

import (
	"bytes"
	"errors"
	"net/url"
	"nippo/admin/adminprops"
	"nippo/cognito"
	"nippo/schema"

	"github.com/google/uuid"
)

func Invite(props *adminprops.Props, form url.Values) (schema.CleanText, error) {
	userIdStr := form.Get("user")
	if len(userIdStr) == 0 {
		return "", errors.New("invalid form")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return "", err
	}

	invitedUser, err := props.Db().ReadUserEmailOrgId(userId[:])
	if err != nil {
		return "", err
	}
	if !bytes.Equal(invitedUser.OrgId, props.OrgId()) {
		return "", errors.New("wrong org")
	}
	if len(invitedUser.Email) == 0 {
		return "", errors.New("user email is empty")
	}

	err = cognito.Invite(props, invitedUser.Email)
	if err != nil {
		return "", err
	}

	return invitedUser.Email, nil
}
