package database

import "time"

type Database interface {
	CreateOrg(name string) error
	CreateUser(name string) error
	CreateMorningCall(date time.Time) error
	CreateEveningCall(date time.Time) error
	CreateDailyReport(date time.Time) error
}
