package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti TiDbTx) CreateMorningCall(
	userId []uint8,
	name schema.CleanText,
	carNum schema.CleanText,
	photo string,
) ([]uint8, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}
	if len(photo) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected photo length")
	}
	if len(name) <= 1 || len(name) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected name length")
	}
	if len(carNum) == 0 || len(carNum) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected car number length")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.MorningCall{
		Id:              id[:],
		Caller:          name,
		Name:            name,
		CarNumber:       carNum,
		Method:          schema.METHOD_PERSON,
		UsingAlcChecker: true,
		AlcCheck:        true,
		AlcPhoto:        photo,
	}
	user := schema.User{}
	res := ti.orm.Select("id").Take(&user, "id = ?", userId)
	if res.Error != nil {
		return nil, res.Error
	}
	err = ti.orm.Model(&user).Association("MorningCalls").Append(&call)
	if err != nil {
		return nil, err
	}

	return call.Id, nil
}

func (ti TiDbTx) UpdateMCallHealth(callId []uint8) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}

	call := schema.MorningCall{}
	res := ti.orm.Model(&call).Where("id = ?", callId).Update("health_check", true)
	return res.Error
}

func (ti TiDbTx) UpdateMCallCarCheck(callId []uint8) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}

	call := schema.MorningCall{}
	res := ti.orm.Model(&call).Where("id = ?", callId).Update("car_check", true)
	return res.Error
}

func (ti TiDbTx) UpdateMCallNote(callId []uint8, note schema.CleanText) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}
	if len(note) > schema.MAX_LONG_TEXT_LEN {
		return errors.New("unexpected note length")
	}

	call := schema.MorningCall{}
	res := ti.orm.Model(&call).Where("id = ?", callId).Update("note", note)
	return res.Error
}

func (ti TiDbTx) CreateEveningCall(
	userId []uint8,
	name schema.CleanText,
	carNum schema.CleanText,
	photo string,
) ([]uint8, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}
	if len(photo) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected photo length")
	}
	if len(name) <= 1 || len(name) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected name length")
	}
	if len(carNum) == 0 || len(carNum) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected car number length")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.EveningCall{
		Id:              id[:],
		Caller:          name,
		Name:            name,
		CarNumber:       carNum,
		Method:          schema.METHOD_PERSON,
		UsingAlcChecker: true,
		AlcCheck:        true,
		AlcPhoto:        photo,
	}
	user := schema.User{}
	res := ti.orm.Select("id").Take(&user, "id = ?", userId)
	if res.Error != nil {
		return nil, res.Error
	}
	err = ti.orm.Model(&user).Association("EveningCalls").Append(&call)
	if err != nil {
		return nil, err
	}

	return call.Id, nil
}

func (ti TiDbTx) UpdateECallNote(callId []uint8, note schema.CleanText) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}
	if len(note) > schema.MAX_LONG_TEXT_LEN {
		return errors.New("unexpected note length")
	}

	call := schema.EveningCall{}
	res := ti.orm.Model(&call).Where("id = ?", callId).Update("note", note)
	return res.Error
}
