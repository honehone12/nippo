package database

import (
	"nippo/schema"
	"time"
)

type Database interface {
	CreateOrg(id [16]uint8) (*schema.Org, error)
	CreateUser(
		name string,
		lineId string,
		orgIdStr string,
	) (*schema.User, error)
	CreateMorningCall(
		date time.Time,
		userLineId string,
	) (*schema.MorningCall, error)
	CreateEveningCall(
		date time.Time,
		userLineId string,
	) (*schema.EveningCall, error)
	CreateDailyReport(
		date time.Time,
		userLineId string,
	) (*schema.DailyReport, error)
}
