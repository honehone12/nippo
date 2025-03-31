package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti *TiDb) CreateUser(
	orgIdStr string,
	name string,
	lineId string,
) ([]uint8, error) {
	if len(name) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected name length")
	}
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
		Name:   name,
		LineId: lineId,
	}
	if err := ti.orm.Model(&schema.Org{
		Id: orgId[:],
	}).Association("Users").Append(&user); err != nil {
		return nil, err
	}

	return user.Id, nil
}

func (ti *TiDb) UpdateUserCarNumber(lineId string, carNumber string) error {
	if len(lineId) != schema.LINE_ID_LEN {
		return errors.New("unexpected line id length")
	}
	if len(carNumber) > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected car number length")
	}

	user := schema.User{}
	res := ti.orm.Model(&user).Where("line_id = ?", lineId).Update("CarNumber", carNumber)
	return res.Error
}

func (ti *TiDb) UpdateUserNextState(lineId string, nextState schema.StateIdx) error {
	if len(lineId) != schema.LINE_ID_LEN {
		return errors.New("unexpected line id length")
	}

	user := schema.User{}
	res := ti.orm.Model(&user).Where("line_id = ?", lineId).Update("NextState", nextState)
	return res.Error
}

func (ti *TiDb) UpdateUserCurrentDuty(lineId string, reportId []uint8) error {
	if len(lineId) != schema.LINE_ID_LEN {
		return errors.New("unexpected line id length")
	}
	if len(reportId) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}

	user := schema.User{}
	res := ti.orm.Model(&user).Where("line_id = ?", lineId).Update("CurrentDuty", reportId)
	return res.Error
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
