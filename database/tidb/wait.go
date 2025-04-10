package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti TiDbTx) CreateWaiting(
	repoId []uint8,
	label schema.CleanText,
	address schema.CleanText,
	latitude float64,
	longitude float64,
) ([]uint8, error) {
	if len(repoId) != schema.ID_LEN {
		return nil, errors.New("unexpected daily report id length")
	}
	if len(label) > schema.MAX_SHORT_TEXT_LEN || len(address) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected text length")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	rest := schema.Waiting{
		Id:        id[:],
		Label:     label,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	}
	repo := schema.DailyReport{Id: repoId}
	err = ti.orm.Model(&repo).Association("Waitings").Append(&rest)
	if err != nil {
		return nil, err
	}

	return rest.Id, nil
}

func (ti TiDb) ReadWaitings(repoId []uint8) ([]schema.Waiting, error) {
	if len(repoId) != schema.ID_LEN {
		return nil, errors.New("unexpected report id length")
	}

	repo := schema.DailyReport{Id: repoId}
	var waits []schema.Waiting
	err := ti.orm.Model(&repo).Association("Waitings").Find(&waits)
	if err != nil {
		return nil, err
	}

	return waits, nil
}

func (ti TiDbTx) UpdateWaitNote(waitId []uint8, note schema.CleanText) error {
	if len(waitId) != schema.ID_LEN {
		return errors.New("unexpected wait id length")
	}
	if len(note) > schema.MAX_LONG_TEXT_LEN {
		return errors.New("unexpected note length")
	}

	wait := schema.Waiting{Id: waitId}
	res := ti.orm.Model(&wait).Update("note", note)
	return res.Error
}
