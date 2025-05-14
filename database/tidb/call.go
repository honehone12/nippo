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
	photo string,
) ([]uint8, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}
	if len(photo) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected photo length")
	}
	if len := len(name); len == 0 || len > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected name length")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.MorningCall{
		Id:              id[:],
		Caller:          name,
		Name:            name,
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

func (ti TiDb) ReadMorningCall(id []uint8) (*schema.MorningCall, error) {
	if len(id) != schema.ID_LEN {
		return nil, errors.New("unexpexted call id length")
	}

	call := schema.MorningCall{Id: id}
	if res := ti.orm.Take(&call); res.Error != nil {
		return nil, res.Error
	}

	return &call, nil
}

func (ti TiDb) ReadMCallCarNum(id []uint8) (schema.CleanText, error) {
	if len(id) != schema.ID_LEN {
		return "", errors.New("unexpexted call id length")
	}

	call := schema.MorningCall{Id: id}
	if res := ti.orm.Select("car_number").Take(&call); res.Error != nil {
		return "", res.Error
	}

	return call.CarNumber, nil
}

func (ti TiDb) ReadMCallAlcPhoto(id []uint8) (string, error) {
	if len(id) != schema.ID_LEN {
		return "", errors.New("unexpexted call id length")
	}

	call := schema.MorningCall{Id: id}
	if res := ti.orm.Select("alc_photo").Take(&call); res.Error != nil {
		return "", res.Error
	}

	return call.AlcPhoto, nil
}

func (ti TiDb) ReadMorningCalls(
	userId []uint8,
	from time.Time,
	to time.Time,
) ([]schema.MorningCall, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}

	user := schema.User{Id: userId}
	var calls []schema.MorningCall
	err := ti.orm.Model(&user).Where(
		"created_at >= ? AND created_at < ?",
		from,
		to,
	).Association("MorningCalls").Find(&calls)
	if err != nil {
		return nil, err
	}

	return calls, nil
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

func (ti TiDbTx) UpdateMCallCarNumber(callId []uint8, carNumber schema.CleanText) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}
	if len := len(carNumber); len == 0 || len > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected car number length")
	}

	call := schema.MorningCall{Id: callId}
	res := ti.orm.Model(&call).Update("car_number", carNumber)
	return res.Error
}

func (ti TiDbTx) UpdateMCallLicenseExp(callId []uint8, exp time.Time) error {
	if len(callId) != schema.ID_LEN {
		return errors.New("unexpected id length")
	}

	call := schema.MorningCall{Id: callId}
	res := ti.orm.Model(&call).Update("license_expiration", exp)
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
	if len := len(name); len == 0 || len > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected name length")
	}
	if len := len(carNum); len == 0 || len > schema.MAX_SHORT_TEXT_LEN {
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

func (ti TiDb) ReadEveningCall(id []uint8) (*schema.EveningCall, error) {
	if len(id) != schema.ID_LEN {
		return nil, errors.New("unexpexted call id length")
	}

	call := schema.EveningCall{Id: id}
	if res := ti.orm.Take(&call); res.Error != nil {
		return nil, res.Error
	}

	return &call, nil
}

func (ti TiDb) ReadECallAlcPhoto(id []uint8) (string, error) {
	if len(id) != schema.ID_LEN {
		return "", errors.New("unexpexted call id length")
	}

	call := schema.EveningCall{Id: id}
	if res := ti.orm.Select("alc_photo").Take(&call); res.Error != nil {
		return "", res.Error
	}

	return call.AlcPhoto, nil
}

func (ti TiDb) ReadEveningCalls(
	userId []uint8,
	from time.Time,
	to time.Time,
) ([]schema.EveningCall, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}

	user := schema.User{Id: userId}
	var calls []schema.EveningCall
	err := ti.orm.Model(&user).Where(
		"created_at >= ? AND created_at < ?",
		from,
		to,
	).Association("EveningCalls").Find(&calls)
	if err != nil {
		return nil, err
	}

	return calls, nil
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
