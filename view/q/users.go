package q

import (
	"encoding/json"
	"errors"
	"nippo/schema"
	"nippo/view/viewprops"

	"gorm.io/gorm"
)

type QUsers struct{}

func (q QUsers) Query(props *viewprops.Props) ([]byte, error) {
	users, err := props.Db().ReadUsers(props.OrgId())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		users = []schema.User{}
	} else if err != nil {
		return nil, err
	}

	return json.Marshal(users)
}
