package tidb

import (
	"errors"
	"nippo/schema"
	"strings"

	"github.com/google/uuid"
)

func (ti TiDbTx) CreateUser(orgIdStr schema.CleanText, lineId schema.CleanText) ([]uint8, error) {
	if len(orgIdStr) != schema.ID_STR_LEN {
		return nil, errors.New("unexpected id length")
	}
	if len(lineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}
	if !strings.HasPrefix(string(lineId), "U") {
		return nil, errors.New("unexpected line id")
	}

	orgId, err := uuid.Parse(string(orgIdStr))
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	user := schema.User{
		Id:     id[:],
		LineId: lineId,
	}
	org := schema.Org{Id: orgId[:]}
	err = ti.orm.Model(&org).Association("Users").Append(&user)
	if err != nil {
		return nil, err
	}

	return user.Id, nil
}

func (ti TiDb) ReadUser(lineId schema.CleanText) (*schema.User, error) {
	if len(lineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}
	if !strings.HasPrefix(string(lineId), "U") {
		return nil, errors.New("unexpected line id")
	}

	user := schema.User{}
	if res := ti.orm.Take(&user, "line_id = ?", lineId); res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (ti TiDb) ReadUsers(orgIdStr schema.CleanText) ([]schema.User, error) {
	if len(orgIdStr) != schema.ID_STR_LEN {
		return nil, errors.New("unexpected id length")
	}

	id, err := uuid.Parse(string(orgIdStr))
	if err != nil {
		return nil, err
	}

	org := schema.Org{Id: id[:]}
	var users []schema.User
	if err := ti.orm.Model(&org).Association("Users").Find(&users); err != nil {
		return nil, err
	}

	return users, err
}

func (ti TiDbTx) UpdateUserName(userId []uint8, name schema.CleanText) error {
	if len(userId) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}
	if len(name) <= 1 || len(name) > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected name length")
	}

	user := schema.User{Id: userId}
	res := ti.orm.Model(&user).Update("name", name)
	return res.Error
}

func (ti TiDbTx) UpdateUserCarNumber(userId []uint8, carNumber schema.CleanText) error {
	if len(userId) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}
	if len(carNumber) == 0 || len(carNumber) > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected car number length")
	}

	user := schema.User{Id: userId}
	res := ti.orm.Model(&user).Update("car_number", carNumber)
	return res.Error
}

func (ti TiDbTx) UpdateUserNextState(userId []uint8, nextState schema.StateIdx) error {
	if len(userId) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}

	user := schema.User{Id: userId}
	res := ti.orm.Model(&user).Update("next_state", nextState)
	return res.Error
}

func (ti TiDbTx) UpdateUserCurrentOpe(userId []uint8, opeId []uint8) error {
	if len(userId) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}
	if opeId != nil && len(opeId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}

	user := schema.User{Id: userId}
	res := ti.orm.Model(&user).Update("current_operation", opeId)
	return res.Error
}

func (ti TiDbTx) UpdateUserCurrentDuty(userId []uint8, repoId []uint8) error {
	if len(userId) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}
	if repoId != nil && len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}

	user := schema.User{Id: userId}
	res := ti.orm.Model(&user).Update("current_duty", repoId)
	return res.Error
}
