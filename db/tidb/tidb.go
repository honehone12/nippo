package tidb

import (
	"nippo/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TiDb struct {
	orm *gorm.DB
}

func Open(tiUri string) (TiDb, error) {
	orm, err := gorm.Open(mysql.Open(tiUri))
	if err != nil {
		return TiDb{nil}, err
	}

	return TiDb{orm}, nil
}

func (ti *TiDb) AutoMigrate() error {
	return ti.orm.AutoMigrate(
		&schema.Org{},
		&schema.User{},
		&schema.Call{},
		&schema.Report{},
	)
}
