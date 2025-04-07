package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti TiDbTx) CreateOrg(name schema.CleanText) ([]uint8, error) {
	if len(name) <= 1 || len(name) > schema.MAX_SHORT_TEXT_LEN {
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

func (ti TiDb) ReadOrg(idStr string) (*schema.Org, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	org := schema.Org{Id: id[:]}
	if res := ti.orm.Take(&org); res.Error != nil {
		return nil, res.Error
	}

	return &org, nil
}

func (ti TiDbTx) UpdateOrgLineId(idStr string, newLineId schema.CleanText) error {
	if len(newLineId) != schema.LINE_ID_LEN {
		return errors.New("unexpected line id length")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}

	org := schema.Org{Id: id[:]}
	res := ti.orm.Model(&org).Update("line_id", newLineId)
	return res.Error
}
