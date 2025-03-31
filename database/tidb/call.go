package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti *TiDb) CreateMorningCall(userLineId string) ([]uint8, error) {
	if len(userLineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	user := schema.User{}
	if res := ti.orm.Select("id").Take(&user, "line_id = ?", userLineId); res.Error != nil {
		return nil, res.Error
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.MorningCall{Id: id[:]}
	if err := ti.orm.Model(&user).Association("MorningCalls").Append(&call); err != nil {
		return nil, err
	}

	return call.Id, nil
}

func (ti *TiDb) CreateEveningCall(userLineId string) ([]uint8, error) {
	if len(userLineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	user := schema.User{}
	if res := ti.orm.Select("id").Take(&user, "line_id = ?", userLineId); res.Error != nil {
		return nil, res.Error
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.EveningCall{Id: id[:]}
	if err := ti.orm.Model(&user).Association("EveningCalls").Append(&call); err != nil {
		return nil, err
	}

	return call.Id, nil
}
