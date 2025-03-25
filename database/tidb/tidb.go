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
	name string,
	lineId string,
	adminLineId string,
) (*schema.Org, error) {
	if len(name) > schema.MAX_NAME_LEN {
		return nil, errors.New("unexpected name length")
	}
	if len(lineId) != schema.LINE_ID_LEN || len(adminLineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	org := schema.Org{
		Id:          id[:],
		Name:        name,
		LineId:      lineId,
		AdminLineId: adminLineId,
	}

	err = ti.orm.Create(&org).Error
	if err != nil {
		return nil, err
	}

	return &org, nil
}

func (ti *TiDb) CreateUser(
	name string,
	lineId string,
	orgLineId string,
) (*schema.User, error) {
	if len(name) > schema.MAX_NAME_LEN {
		return nil, errors.New("unexpected name length")
	}
	if len(lineId) != schema.LINE_ID_LEN || len(orgLineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	org := schema.Org{}
	err = ti.orm.
		Select("id").
		Where("line_id = ?", orgLineId).
		First(&org).Error
	if err != nil {
		return nil, err
	}

	user := schema.User{
		Id:     id[:],
		Name:   name,
		LineId: lineId,
	}
	err = ti.orm.
		Model(&org).
		Association("Users").
		Append(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ti *TiDb) CreateMorningCall(
	date time.Time,
	userLineId string,
) (*schema.MorningCall, error) {
	if len(userLineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	user := schema.User{}
	err := ti.orm.
		Select("id").
		Where("line_id = ?", userLineId).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.MorningCall{
		Id:   id[:],
		Date: date,
	}
	err = ti.orm.
		Model(&user).
		Association("MorningCalls").
		Append(&call)
	if err != nil {
		return nil, err
	}

	return &call, nil
}

func (ti *TiDb) CreateEveningCall(
	date time.Time,
	userLineId string,
) (*schema.EveningCall, error) {
	if len(userLineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	user := schema.User{}
	err := ti.orm.
		Select("id").
		Where("line_id = ?", userLineId).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.EveningCall{
		Id:   id[:],
		Date: date,
	}
	err = ti.orm.
		Model(&user).
		Association("EveningCalls").
		Append(&call)
	if err != nil {
		return nil, err
	}

	return &call, nil
}

func (ti *TiDb) CreateDailyReport(
	date time.Time,
	userLineId string,
) (*schema.DailyReport, error) {
	if len(userLineId) != schema.LINE_ID_LEN {
		return nil, errors.New("unexpected line id length")
	}

	user := schema.User{}
	err := ti.orm.
		Select("id").
		Where("line_id = ?", userLineId).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.DailyReport{
		Id:   id[:],
		Date: date,
	}
	err = ti.orm.
		Model(&user).
		Association("DailyReports").
		Append(&call)
	if err != nil {
		return nil, err
	}

	return &call, nil
}
