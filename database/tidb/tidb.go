package tidb

import (
	"errors"
	"nippo/schema"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TiDb struct {
	orm *gorm.DB
}

func Open(tiUri string) (*TiDb, error) {
	orm, err := gorm.Open(mysql.Open(tiUri))
	if err != nil {
		return nil, err
	}

	return &TiDb{orm}, nil
}

func (ti *TiDb) AutoMigrate() error {
	return ti.orm.AutoMigrate(
		&schema.Org{},
		&schema.User{},
		&schema.MorningCall{},
		&schema.EveningCall{},
		&schema.DailyReport{},
	)
}

func (ti *TiDb) CreateOrg(
	id [16]uint8,
	name string,
	lineId string,
) (*schema.Org, error) {
	if len(name) > schema.MAX_NAME_LEN {
		return nil, errors.New("unexpected name length")
	}
	if len(lineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	org := schema.Org{
		Id:     id[:],
		Name:   name,
		LineId: lineId,
	}

	if res := ti.orm.Create(&org); res.Error != nil {
		return nil, res.Error
	}

	return &org, nil
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

func (ti *TiDb) CreateUser(
	orgIdStr string,
	name string,
	lineId string,
) (*schema.User, error) {
	if len(name) > schema.MAX_NAME_LEN {
		return nil, errors.New("unexpected name length")
	}
	if len(lineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}
	if len(orgIdStr) != schema.ID_STR_LEN {
		return nil, errors.New("unexpected id length")
	}

	orgId, err := uuid.Parse(orgIdStr)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	user := schema.User{
		Id:     id[:],
		Name:   name,
		LineId: lineId,
	}
	if err := ti.orm.Model(&schema.Org{
		Id: orgId[:],
	}).Association("Users").Append(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ti *TiDb) CreateMorningCall(
	userLineId string,
	date time.Time,
) (*schema.MorningCall, error) {
	if len(userLineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	user := schema.User{}
	if res := ti.orm.Select("id").Take(&user, "line_id = ?", userLineId); res.Error != nil {
		return nil, res.Error
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.MorningCall{
		Id:   id[:],
		Date: date,
	}
	if err := ti.orm.Model(&user).Association("MorningCalls").Append(&call); err != nil {
		return nil, err
	}

	return &call, nil
}

func (ti *TiDb) CreateEveningCall(
	userLineId string,
	date time.Time,
) (*schema.EveningCall, error) {
	if len(userLineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	user := schema.User{}
	if res := ti.orm.Select("id").Take(&user, "line_id = ?", userLineId); res.Error != nil {
		return nil, res.Error
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.EveningCall{
		Id:   id[:],
		Date: date,
	}
	if err := ti.orm.Model(&user).Association("EveningCalls").Append(&call); err != nil {
		return nil, err
	}

	return &call, nil
}

func (ti *TiDb) CreateDailyReport(
	userLineId string,
	date time.Time,
) (*schema.DailyReport, error) {
	if len(userLineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	user := schema.User{}
	if res := ti.orm.Select("id").Take(&user, "line_id = ?", userLineId); res.Error != nil {
		return nil, res.Error
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.DailyReport{
		Id:   id[:],
		Date: date,
	}
	if err := ti.orm.Model(&user).Association("DailyReports").Append(&call); err != nil {
		return nil, err
	}

	return &call, nil
}
