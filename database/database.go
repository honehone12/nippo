package database

import (
	"nippo/schema"
	"time"
)

type Database interface {
	CreateOrg(name string) ([]uint8, error)
	UpdateOrgLineId(idStr string, newLineId string) error
	ReadOrg(idStr string) (*schema.Org, error)

	CreateUser(
		orgIdStr string,
		name string,
		lineId string,
	) ([]uint8, error)
	ReadUser(userLineId string) (*schema.User, error)

	CreateMorningCall(userLineId string, date time.Time) ([]uint8, error)

	CreateEveningCall(userLineId string, date time.Time) ([]uint8, error)

	CreateDailyReport(userLineId string, date time.Time) ([]uint8, error)
}
