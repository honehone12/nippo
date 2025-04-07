package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti TiDbTx) CreateLocation(
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

	repo := schema.DailyReport{Id: repoId}
	loc := schema.Location{
		Id:        id[:],
		Label:     label,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	}
	if err := ti.orm.Model(&repo).Association("Locations").Append(&loc); err != nil {
		return nil, err
	}

	return loc.Id, nil
}

func (ti TiDbTx) UpdateLocNote(locId []uint8, shortNote schema.CleanText) error {
	if len(locId) != schema.ID_LEN {
		return errors.New("unexpected location id length")
	}
	if len(shortNote) > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected note length")
	}

	loc := schema.Location{Id: locId}
	res := ti.orm.Model(&loc).Update("short_note", shortNote)
	return res.Error
}
