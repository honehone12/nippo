package q

import (
	"encoding/json"
	"nippo/view/viewprops"
)

type QUsers struct{}

func (q QUsers) Query(props *viewprops.Props) ([]byte, error) {
	users, err := props.Db().ReadUsers(props.OrgId())
	if err != nil {
		return nil, err
	}

	return json.Marshal(users)
}
