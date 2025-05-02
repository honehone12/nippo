package q

import (
	"encoding/json"
	"errors"
	"nippo/cognito"
	"nippo/schema"
	"nippo/view/viewprops"

	"gorm.io/gorm"
)

type QUsers struct{}

func (q QUsers) Query(props *viewprops.Props, email schema.CleanText) ([]byte, error) {
	db := props.Db()

	orgEmail, err := db.ReadOrgEmail(props.OrgId())
	if err != nil {
		return nil, err
	}
	admin := orgEmail == email

	read, err := db.ReadUsersMini(props.OrgId())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		users := schema.Users{
			Admin:      admin,
			Users:      []schema.UserMini{},
			Invitables: []schema.UserMini{},
		}
		return json.Marshal(users)
	} else if err != nil {
		return nil, err
	}

	users := schema.Users{
		Admin:      admin,
		Users:      make([]schema.UserMini, 0, len(read)),
		Invitables: []schema.UserMini{},
	}
	for _, u := range read {
		user := schema.UserMini{
			Id:        u.Id,
			CreatedAt: u.CreatedAt,
			Name:      u.Name,
		}
		users.Users = append(users.Users, user)

		if admin {
			invite := cognito.Pooled(props, u.Email)
			if invite {
				users.Invitables = append(users.Invitables, user)
			}
		}
	}

	return json.Marshal(users)
}
