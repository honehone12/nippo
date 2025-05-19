package tidb

import (
	"errors"
	"nippo/schema"
	"strings"

	"github.com/google/uuid"
)

func (ti TiDbTx) CreateOrg(name schema.CleanText, email schema.CleanText) ([]uint8, error) {
	if len := len(name); len == 0 || len > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected name length")
	}
	if len := len(email); len == 0 || len > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected email length")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	org := schema.Org{
		Id:    id[:],
		Name:  name,
		Email: email,
	}

	if res := ti.orm.Create(&org); res.Error != nil {
		return nil, res.Error
	}

	return org.Id, nil
}

func (ti TiDb) ReadOrg(id []uint8) (*schema.Org, error) {
	if len(id) != schema.ID_LEN {
		return nil, errors.New("unexpected org id length")
	}

	org := &schema.Org{Id: id}
	if res := ti.orm.Take(org); res.Error != nil {
		return nil, res.Error
	}

	return org, nil
}

func (ti TiDb) ReadOrgLineId(id []uint8) (schema.CleanText, error) {
	if len(id) != schema.ID_LEN {
		return "", errors.New("unexpected org id length")
	}

	org := schema.Org{Id: id}
	if res := ti.orm.Select("line_id").Take(&org); res.Error != nil {
		return "", res.Error
	}

	return org.LineId, nil
}

func (ti TiDb) ReadOrgEmail(id []uint8) (schema.CleanText, error) {
	if len(id) != schema.ID_LEN {
		return "", errors.New("unexpected org id length")
	}

	org := schema.Org{Id: id}
	if res := ti.orm.Select("email").Take(&org); res.Error != nil {
		return "", res.Error
	}

	return org.Email, nil
}

func (ti TiDbTx) UpdateOrgLineId(id []uint8, newLineId schema.CleanText) error {
	if len(id) != schema.ID_LEN {
		return errors.New("unexpected org id length")
	}
	if len(newLineId) != schema.LINE_ID_LEN {
		return errors.New("unexpected line id length")
	}
	if !strings.HasPrefix(string(newLineId), "U") {
		return errors.New("unexpected line id")
	}

	org := schema.Org{Id: id}
	res := ti.orm.Model(&org).Update("line_id", newLineId)
	return res.Error
}
