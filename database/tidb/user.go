package tidb

import (
	"errors"
	"nippo/schema"
	"strings"

	"github.com/google/uuid"
)

func (ti TiDbTx) CreateUser(orgId []uint8, lineId schema.CleanText) ([]uint8, error) {
	if len(orgId) != schema.ID_LEN {
		return nil, errors.New("unexpected org id length")
	}
	if len(lineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}
	if !strings.HasPrefix(string(lineId), "U") {
		return nil, errors.New("unexpected line id")
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

func (ti TiDb) ReadUserEmailOrgId(id []uint8) (*schema.User, error) {
	if len(id) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}

	user := schema.User{Id: id}
	if res := ti.orm.Select("email", "org_id").Take(&user); res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (ti TiDb) ReadUser_Line(lineId schema.CleanText) (*schema.User, error) {
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

func (ti TiDb) ReadUserOrgId(id []uint8) ([]uint8, error) {
	if len(id) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}

	user := schema.User{Id: id}
	if res := ti.orm.Select("org_id").Take(&user); res.Error != nil {
		return nil, res.Error
	}

	return user.OrgId, nil
}

func (ti TiDb) ReadUserOrgId_Email(email schema.CleanText) ([]uint8, error) {
	if len := len(email); len == 0 || len > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected line id length")
	}

	user := schema.User{}
	if res := ti.orm.Select("org_id").Take(&user, "email = ?", email); res.Error != nil {
		return nil, res.Error
	}

	return user.OrgId, nil
}

func (ti TiDb) ReadUserBasic_Email(email schema.CleanText) (*schema.User, error) {
	if len := len(email); len == 0 || len > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected line id length")
	}

	user := schema.User{}
	if res := ti.orm.Select(
		"id",
		"created_at",
		"name",
		"sub_admin",
		"org_id",
	).Take(&user, "email = ?", email); res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (ti TiDb) ReadUserEmail(id []uint8) (schema.CleanText, error) {
	if len(id) != schema.ID_LEN {
		return "", errors.New("unexpected user id length")
	}

	user := schema.User{Id: id}
	if res := ti.orm.Select("email").Take(&user); res.Error != nil {
		return "", res.Error
	}

	return user.Email, nil
}

func (ti TiDb) ReadUsersMini(orgId []uint8) ([]schema.User, error) {
	if len(orgId) != schema.ID_LEN {
		return nil, errors.New("unexpected org id length")
	}

	org := schema.Org{Id: orgId}
	var users []schema.User
	err := ti.orm.Model(&org).Select(
		"id",
		"created_at",
		"name",
		"sub_admin",
	).Association("Users").Find(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ti TiDbTx) UpdateUserName(userId []uint8, name schema.CleanText) error {
	if len(userId) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}
	if len := len(name); len == 0 || len > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected name length")
	}

	user := schema.User{Id: userId}
	res := ti.orm.Model(&user).Update("name", name)
	return res.Error
}

func (ti TiDbTx) UpdateUserSubAdmin(userId []uint8, subAdmin bool) error {
	if len(userId) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}

	user := schema.User{Id: userId}
	res := ti.orm.Model(&user).Update("sub_admin", subAdmin)
	return res.Error
}

func (ti TiDbTx) UpdateUserEmail(userId []uint8, email schema.CleanText) error {
	if len(userId) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}
	if len := len(email); len == 0 || len > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected email length")
	}

	user := schema.User{Id: userId}
	res := ti.orm.Model(&user).Update("email", email)
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
