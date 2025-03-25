package database

import (
	"nippo/schema"
	"time"
)

type Database interface {
	CreateOrg(
		name string,
		lineId string,
		adminLineId string,
	) (*schema.Org, error)
	CreateUser(
		name string,
		lineId string,
		orgLineId string,
	) (*schema.User, error)
	CreateMorningCall(date time.Time, userLineId string) error
	CreateEveningCall(date time.Time, userLineId string) error
	CreateDailyReport(date time.Time, userLineId string) error
}
