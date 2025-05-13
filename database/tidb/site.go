package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti TiDbTx) CreateSite(repoId []uint8, name schema.CleanText) ([]uint8, error) {
	if len(repoId) != schema.ID_LEN {
		return nil, errors.New("unexpected daily report id length")
	}
	if len := len(name); len == 0 || len > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected name length")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	site := schema.Site{
		Id:   id[:],
		Name: name,
	}
	repo := schema.DailyReport{Id: repoId}
	err = ti.orm.Model(&repo).Association("Sites").Append(&site)
	if err != nil {
		return nil, err
	}

	return site.Id, nil
}

func (ti TiDb) ReadSites(repoId []uint8) ([]schema.Site, error) {
	if len(repoId) != schema.ID_LEN {
		return nil, errors.New("unexpected report id length")
	}

	repo := schema.DailyReport{Id: repoId}
	var sites []schema.Site
	err := ti.orm.Model(&repo).Association("Sites").Find(&sites)
	if err != nil {
		return nil, err
	}

	return sites, nil
}

func (ti TiDbTx) UpdateSiteNote(siteId []uint8, shortNote schema.CleanText) error {
	if len(siteId) != schema.ID_LEN {
		return errors.New("unexpected site id length")
	}
	if len(shortNote) > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected note length")
	}

	site := schema.Site{Id: siteId}
	res := ti.orm.Model(&site).Update("short_note", shortNote)
	return res.Error
}
