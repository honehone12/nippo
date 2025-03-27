package database

import (
	"nippo/schema"
	"time"
)

type Database interface {
	CreateOrg(
		name string,
		lineId string,
	) (*schema.Org, error)
	ReadOrg(idStr string) (*schema.Org, error)

	CreateUser(
		orgIdStr string,
		name string,
		lineId string,
	) (*schema.User, error)
	ReadUser(userLineId string) (*schema.User, error)

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
