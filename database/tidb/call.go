package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti *TiDb) CreateMorningCall(userId []uint8, photo string) ([]uint8, error) {
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
		Photo:           photo,
	}
	if err := ti.orm.Model(&user).Association("MorningCalls").Append(&call); err != nil {
		return nil, err
	}

	return call.Id, nil
}

func (ti *TiDb) UpdateMoCallHealth(userId []uint8) error {
	if len(userId) != schema.ID_LEN {
		return errors.New("unexpected user id length")
	}

	user := schema.User{}
	if res := ti.orm.Select("current_call").Take(&user, "id = ?", userId); res.Error != nil {
		return res.Error
	}

	if len(user.CurrentCall) == 0 {
		return errors.New("user does not have current call")
	}

	call := schema.MorningCall{}
	res := ti.orm.Model(&call).Where("id = ?", user.CurrentCall).Update("health_check", true)
	return res.Error
}

func (ti *TiDb) UpdateMoCallCarCheck(userId []uint8) error {
	if len(userId) != schema.ID_LEN {
		return errors.New("unexpected user id length")
	}

	user := schema.User{}
	if res := ti.orm.Select("current_call").Take(&user, "id = ?", userId); res.Error != nil {
		return res.Error
	}

	if len(user.CurrentCall) == 0 {
		return errors.New("user does not have current call")
	}

	call := schema.MorningCall{}
	res := ti.orm.Model(&call).Where("id = ?", user.CurrentCall).Update("car_check", true)
	return res.Error
}

func (ti *TiDb) UpdateMoCallNote(userId []uint8, note string) error {
	if len(userId) != schema.ID_LEN {
		return errors.New("unexpected user id length")
	}
	if len(note) > schema.MAX_LONG_TEXT_LEN {
		return errors.New("unexpected note length")
	}

	user := schema.User{}
	if res := ti.orm.Select("current_call").Take(&user, "id = ?", userId); res.Error != nil {
		return res.Error
	}

	if len(user.CurrentCall) == 0 {
		return errors.New("user does not have current call")
	}

	call := schema.MorningCall{}
	res := ti.orm.Model(&call).Where("id = ?", user.CurrentCall).Update("note", note)
	return res.Error
}

func (ti *TiDb) CreateEveningCall(userId []uint8) ([]uint8, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}

	user := schema.User{}
	if res := ti.orm.Select(
		"id",
		"name",
	).Take(&user, "id = ?", userId); res.Error != nil {
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
