package tidb

import (
	"errors"
	"nippo/schema"
	"time"

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
	user := schema.User{Id: userId}
	err = ti.orm.Model(&user).Association("MorningCalls").Append(&call)
	if err != nil {
		return nil, err
	}

	return call.Id, nil
}

func (ti TiDb) ReadMorningCalls(
	year int,
	from time.Time,
	to time.Time,
) ([]schema.MorningCall, error) {
	panic("not implemented")
}

func (ti TiDbTx) UpdateMCallHealth(callId []uint8) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}

	call := schema.MorningCall{Id: callId}
	res := ti.orm.Model(&call).Update("health_check", true)
	return res.Error
}

func (ti TiDbTx) UpdateMCallCarCheck(callId []uint8) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}

	call := schema.MorningCall{Id: callId}
	res := ti.orm.Model(&call).Update("car_check", true)
	return res.Error
}

func (ti TiDbTx) UpdateMCallNote(callId []uint8, note schema.CleanText) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}
	if len(note) > schema.MAX_LONG_TEXT_LEN {
		return errors.New("unexpected note length")
	}

	call := schema.MorningCall{Id: callId}
	res := ti.orm.Model(&call).Update("note", note)
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
	user := schema.User{Id: userId}
	err = ti.orm.Model(&user).Association("EveningCalls").Append(&call)
	if err != nil {
		return nil, err
	}

	return call.Id, nil
}

func (ti TiDb) ReadEveningCalls(
	year int,
	from time.Time,
	to time.Time,
) ([]schema.EveningCall, error) {
	panic("not implemented")
}

func (ti TiDbTx) UpdateECallNote(callId []uint8, note schema.CleanText) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}
	if len(note) > schema.MAX_LONG_TEXT_LEN {
		return errors.New("unexpected note length")
	}

	call := schema.EveningCall{Id: callId}
	res := ti.orm.Model(&call).Update("note", note)
	return res.Error
}
