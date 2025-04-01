package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti *TiDb) CreateOrg(name string) ([]uint8, error) {
	if len(name) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected name length")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	org := schema.Org{
		Id:   id[:],
		Name: name,
	}

	if res := ti.orm.Create(&org); res.Error != nil {
		return nil, res.Error
	}

	return org.Id, nil
}

func (ti *TiDb) UpdateOrgLineId(idStr string, newLineId string) error {
	if len(newLineId) != schema.LINE_ID_LEN {
		return errors.New("unexpected line id length")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}

	org := schema.Org{}
	res := ti.orm.Model(&org).Where("id = ?", id[:]).Update("line_id", newLineId)
	return res.Error
}

func (ti *TiDb) ReadOrg(idStr string) (*schema.Org, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	org := schema.Org{}
	if res := ti.orm.Take(&org, "id = ?", id[:]); res.Error != nil {
		return nil, res.Error
	}

	return &org, nil
}
