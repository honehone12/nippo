package q

import (
	"bytes"
	"encoding/json"
	"errors"
	"nippo/cognito"
	"nippo/schema"
	"nippo/view/viewprops"

	"github.com/google/uuid"
)

type QInvite struct {
	userId uuid.UUID
}

func (q QInvite) Query(props *viewprops.Props, email schema.CleanText) ([]byte, error) {
	db := props.Db()
	orgId := props.OrgId()

	orgEmail, err := db.ReadOrgEmail(orgId)
	if err != nil {
		return nil, err
	}
	if orgEmail != email {
		return nil, errors.New("not admin")
	}

	invitedUser, err := db.ReadUserEmailOrgId(q.userId[:])
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(invitedUser.OrgId, orgId) {
		return nil, errors.New("wrong org")
	}

	err = cognito.Invite(props, invitedUser.Email)
	if err != nil {
		return nil, err
	}

	invited := schema.Invited{Email: invitedUser.Email}
	return json.Marshal(invited)
}
