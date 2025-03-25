package db

import "time"

type Db interface {
	CreateOrg(name string) error
	CreateUser(name string) error
	CreateMorningCall(date time.Time) error
	CreateEveningCall(date time.Time) error
	CreateDailyReport(date time.Time) error
}
