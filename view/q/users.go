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

func (q QUsers) Query(props *viewprops.Props, viewer *viewprops.Viewer) ([]byte, error) {
	admin := viewer.Admin()
	users := schema.Users{
		Admin:      admin,
		Users:      []schema.UserMini{},
		Invitables: []schema.UserMini{},
	}

	if viewer.Privileged() {
		basics, err := props.Db().ReadUsersBasic(props.OrgId())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return json.Marshal(users)
		} else if err != nil {
			return nil, err
		}

		for _, u := range basics {
			user := schema.UserMini{
				Id:        u.Id,
				CreatedAt: u.CreatedAt,
				Name:      u.Name,
			}
			users.Users = append(users.Users, user)

			if !admin {
				continue
			}

			if len(u.Email) == 0 {
				continue
			}

			exists, err := cognito.Exists(props, u.Email)
			if err != nil {
				return nil, err
			}
			if !exists {
				users.Invitables = append(users.Invitables, user)
			}
		}
	} else {
		users.Users = append(users.Users, schema.UserMini{
			Id:        viewer.Id(),
			CreatedAt: viewer.CreatedAt(),
			Name:      viewer.Name(),
		})
	}

	return json.Marshal(users)
}
