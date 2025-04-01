package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti *TiDb) CreateDailyReport(userId []uint8, morningId []uint8) ([]uint8, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}
	if len(morningId) != schema.ID_LEN {
		return nil, errors.New("unexpected morning id len")
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

	call := schema.DailyReport{
		Id:            id[:],
		MorningCallId: morningId,
		Name:          user.Name,
		CarNumber:     user.CarNumber,
	}
	if err := ti.orm.Model(&user).Association("DailyReports").Append(&call); err != nil {
		return nil, err
	}

	return call.Id, nil
}
