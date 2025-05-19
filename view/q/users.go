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
		Admin: admin,
		Users: []schema.UserInfo{},
	}

	if viewer.Privileged() {
		userMinis, err := props.Db().ReadUsersMini(props.OrgId())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return json.Marshal(users)
		} else if err != nil {
			return nil, err
		}

		for _, u := range userMinis {
			info := schema.UserInfo{
				User: schema.UserMini{
					Id:        u.Id,
					CreatedAt: u.CreatedAt,
					Name:      u.Name,
				},
			}

			if admin && len(u.Email) > 0 {
				exists, err := cognito.Exists(props, u.Email)
				if err != nil {
					return nil, err
				}

				info.Invitable = !exists
				info.Promotable = exists && !u.SubAdmin
			}

			users.Users = append(users.Users, info)
		}
	} else {
		users.Users = append(users.Users, schema.UserInfo{
			User: schema.UserMini{
				Id:        viewer.Id(),
				CreatedAt: viewer.CreatedAt(),
				Name:      viewer.Name(),
			},
		})
	}

	return json.Marshal(users)
}
