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
