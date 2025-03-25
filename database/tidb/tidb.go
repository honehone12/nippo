package tidb

import (
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

func (ti *TiDb) CreateOrg(name string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	return ti.orm.Create(&schema.Org{
		Id:   id[:],
		Name: name,
	}).Error
}

func (ti *TiDb) CreateUser(name string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	return ti.orm.Create(&schema.User{
		Id:   id[:],
		Name: name,
	}).Error
}

func (ti *TiDb) CreateMorningCall(date time.Time) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	return ti.orm.Create(&schema.MorningCall{
		Id:   id[:],
		Date: date,
	}).Error
}

func (ti *TiDb) CreateEveningCall(date time.Time) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	return ti.orm.Create(&schema.EveningCall{
		Id:   id[:],
		Date: date,
	}).Error
}

func (ti *TiDb) CreateDailyReport(date time.Time) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	return ti.orm.Create(&schema.DailyReport{
		Id:   id[:],
		Date: date,
	}).Error
}
