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

	loc := schema.Location{
		Id:        id[:],
		Label:     label,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	}
	repo := schema.DailyReport{Id: repoId}
	err = ti.orm.Model(&repo).Association("Locations").Append(&loc)
	if err != nil {
		return nil, err
	}

	return loc.Id, nil
}

func (ti TiDb) ReadLocations(repoIdStr schema.CleanText) ([]schema.Location, error) {
	if len(repoIdStr) != schema.ID_STR_LEN {
		return nil, errors.New("unexpected location id length")
	}

	repoId, err := uuid.Parse(string(repoIdStr))
	if err != nil {
		return nil, err
	}

	repo := schema.DailyReport{Id: repoId[:]}
	var locs []schema.Location
	err = ti.orm.Model(&repo).Association("Locations").Find(&locs)
	if err != nil {
		return nil, err
	}

	return locs, nil
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
