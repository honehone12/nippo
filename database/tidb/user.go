package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti *TiDb) CreateUser(orgIdStr string, lineId string) ([]uint8, error) {
	if len(lineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}
	if len(orgIdStr) != schema.ID_STR_LEN {
		return nil, errors.New("unexpected id length")
	}

	orgId, err := uuid.Parse(orgIdStr)
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
	if err := ti.orm.Model(&schema.Org{
		Id: orgId[:],
	}).Association("Users").Append(&user); err != nil {
		return nil, err
	}

	return user.Id, nil
}

func (ti *TiDb) ReadUser(lineId string) (*schema.User, error) {
	if len(lineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	user := schema.User{}
	if res := ti.orm.Take(&user, "line_id = ?", lineId); res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (ti *TiDb) UpdateUserName(id []uint8, name schema.CleanText) error {
	if len(id) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}
	if len(name) > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected name length")
	}

	user := schema.User{}
	res := ti.orm.Model(&user).Where("id = ?", id).Update("name", name)
	return res.Error
}

func (ti *TiDb) UpdateUserCarNumber(id []uint8, carNumber schema.CleanText) error {
	if len(id) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}
	if len(carNumber) > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected car number length")
	}

	user := schema.User{}
	res := ti.orm.Model(&user).Where("id = ?", id).Update("car_number", carNumber)
	return res.Error
}

func (ti *TiDb) UpdateUserNextState(id []uint8, nextState schema.StateIdx) error {
	if len(id) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}

	user := schema.User{}
	res := ti.orm.Model(&user).Where("id = ?", id).Update("next_state", nextState)
	return res.Error
}

func (ti *TiDb) UpdateUserCurrentCall(id []uint8, callId []uint8) error {
	if len(id) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}
	if callId != nil && len(callId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}

	user := schema.User{}
	res := ti.orm.Model(&user).Where("id = ?", id).Update("current_call", callId)
	return res.Error
}

func (ti *TiDb) UpdateUserCurrentDuty(id []uint8, reportId []uint8) error {
	if len(id) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}
	if reportId != nil && len(reportId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}

	user := schema.User{}
	res := ti.orm.Model(&user).Where("id = ?", id).Update("current_duty", reportId)
	return res.Error
}
