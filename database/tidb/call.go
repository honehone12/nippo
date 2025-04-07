package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti TiDbTx) CreateMorningCall(userId []uint8, photo string) ([]uint8, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}
	if len(photo) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected photo length")
	}

	user := schema.User{}
	if res := ti.orm.Select(
		"id",
		"name",
		"car_number",
	).Take(&user, "id = ?", userId); res.Error != nil {
		return nil, res.Error
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.MorningCall{
		Id:              id[:],
		Caller:          user.Name,
		Name:            user.Name,
		CarNumber:       user.CarNumber,
		Method:          schema.METHOD_PERSON,
		UsingAlcChecker: true,
		AlcCheck:        true,
		AlcPhoto:        photo,
	}
	if err := ti.orm.Model(&user).Association("MorningCalls").Append(&call); err != nil {
		return nil, err
	}

	return call.Id, nil
}

func (ti TiDbTx) UpdateMoCallHealth(callId []uint8) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}

	call := schema.MorningCall{}
	res := ti.orm.Model(&call).Where("id = ?", callId).Update("health_check", true)
	return res.Error
}

func (ti TiDbTx) UpdateMoCallCarCheck(callId []uint8) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}

	call := schema.MorningCall{}
	res := ti.orm.Model(&call).Where("id = ?", callId).Update("car_check", true)
	return res.Error
}

func (ti TiDbTx) UpdateMoCallNote(callId []uint8, note schema.CleanText) error {
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

func (ti TiDbTx) CreateEveningCall(userId []uint8, photo string) ([]uint8, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}
	if len(photo) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected photo length")
	}

	user := schema.User{}
	if res := ti.orm.Select(
		"id",
		"name",
		"car_number",
	).Take(&user, "id = ?", userId); res.Error != nil {
		return nil, res.Error
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.EveningCall{
		Id:              id[:],
		Caller:          user.Name,
		Name:            user.Name,
		CarNumber:       user.CarNumber,
		Method:          schema.METHOD_PERSON,
		UsingAlcChecker: true,
		AlcCheck:        true,
		AlcPhoto:        photo,
	}
	if err := ti.orm.Model(&user).Association("EveningCalls").Append(&call); err != nil {
		return nil, err
	}

	return call.Id, nil
}

func (ti TiDbTx) UpdateEvCallNote(callId []uint8, note schema.CleanText) error {
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
