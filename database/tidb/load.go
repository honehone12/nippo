package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti TiDbTx) CreateLoading(
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

	rest := schema.Loading{
		Id:        id[:],
		Label:     label,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	}
	repo := schema.DailyReport{Id: repoId}
	err = ti.orm.Model(&repo).Association("Loadings").Append(&rest)
	if err != nil {
		return nil, err
	}

	return rest.Id, nil
}

func (ti TiDb) ReadLoadings(repoIdStr schema.CleanText) ([]schema.Loading, error) {
	if len(repoIdStr) != schema.ID_STR_LEN {
		return nil, errors.New("unexpected load id length")
	}

	repoId, err := uuid.Parse(string(repoIdStr))
	if err != nil {
		return nil, err
	}

	repo := schema.DailyReport{Id: repoId[:]}
	var loads []schema.Loading
	err = ti.orm.Model(&repo).Association("Loadings").Find(&loads)
	if err != nil {
		return nil, err
	}

	return loads, nil
}

func (ti TiDbTx) UpdateLoadShipCheck(loadId []uint8, check bool) error {
	if len(loadId) != schema.ID_LEN {
		return errors.New("unexpected wait id length")
	}

	wait := schema.Loading{Id: loadId}
	res := ti.orm.Model(&wait).Update("shipping_check", check)
	return res.Error
}

func (ti TiDbTx) UpdateLoadNote(loadId []uint8, note schema.CleanText) error {
	if len(loadId) != schema.ID_LEN {
		return errors.New("unexpected wait id length")
	}
	if len(note) > schema.MAX_LONG_TEXT_LEN {
		return errors.New("unexpected note length")
	}

	wait := schema.Loading{Id: loadId}
	res := ti.orm.Model(&wait).Update("note", note)
	return res.Error
}
