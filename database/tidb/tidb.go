package tidb

import (
	stdlog "log"
	"nippo/database"
	"nippo/schema"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TiDb struct {
	orm *gorm.DB
}

type TiDbTx struct {
	orm *gorm.DB
}

func now() time.Time {
	return time.Now().UTC()
}

func Open(tiUri string) (TiDb, error) {
	gormLogger := logger.New(
		stdlog.New(os.Stdout, "\r\n", stdlog.LstdFlags),
		logger.Config{
			SlowThreshold:             500 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	orm, err := gorm.Open(mysql.Open(tiUri), &gorm.Config{
		Logger:  gormLogger,
		NowFunc: now,
	})
	if err != nil {
		return TiDb{nil}, err
	}

	return TiDb{orm}, nil
}

func (ti TiDb) Tx(txf func(database.Transaction) error) error {
	return ti.orm.Transaction(func(tx *gorm.DB) error {
		return txf(TiDbTx{tx})
	})
}

func (ti TiDb) Once() database.Transaction {
	return TiDbTx(ti)
}

func (ti TiDb) AutoMigrate() error {
	return ti.orm.AutoMigrate(
		&schema.Org{},
		&schema.User{},
		&schema.MorningCall{},
		&schema.EveningCall{},
		&schema.DailyReport{},
		&schema.Location{},
		&schema.Waiting{},
		&schema.Loading{},
	)
}
