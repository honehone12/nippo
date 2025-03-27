package database

import (
	"nippo/schema"
	"time"
)

type Database interface {
	CreateOrg(
		id [16]uint8,
		name string,
		lineId string,
	) (*schema.Org, error)
	CreateUser(
		orgIdStr string,
		name string,
		lineId string,
	) (*schema.User, error)
	CreateMorningCall(
		userLineId string,
		date time.Time,
	) (*schema.MorningCall, error)
	CreateEveningCall(
		userLineId string,
		date time.Time,
	) (*schema.EveningCall, error)
	CreateDailyReport(
		userLineId string,
		date time.Time,
	) (*schema.DailyReport, error)
}
