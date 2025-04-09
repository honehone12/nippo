package database

import (
	"nippo/schema"
	"time"
)

type Database interface {
	ReadOrg(idStr string) (*schema.Org, error)
	ReadUser(lineId schema.CleanText) (*schema.User, error)
	ReadUsers(orgIdStr string) ([]schema.User, error)
	ReadMorningCalls(
		userIdStr string,
		from time.Time,
		to time.Time,
	) ([]schema.MorningCall, error)
	ReadEveningCalls(
		userIdStr string,
		from time.Time,
		to time.Time,
	) ([]schema.EveningCall, error)
	ReadDailyReports(
		userIdStr string,
		from time.Time,
		to time.Time,
	) ([]schema.DailyReport, error)

	Tx(txf func(Transaction) error) error
	Once() Transaction
}
