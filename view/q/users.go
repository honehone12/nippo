package q

import "nippo/view/viewprops"

type QUsers struct{}

func (q QUsers) Query(props *viewprops.Props) ([]byte, error) {
	users, err := props.Db().ReadUsers(props.OrgId())
	if err != nil {
		return nil, err
	}

}
