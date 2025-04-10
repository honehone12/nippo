package database

import (
	"nippo/schema"
	"time"
)

type Database interface {
	ReadOrg(idStr schema.CleanText) (*schema.Org, error)
	ReadUser(lineId schema.CleanText) (*schema.User, error)
	ReadUsers(orgIdStr schema.CleanText) ([]schema.User, error)
	ReadMorningCalls(
		userIdStr schema.CleanText,
		from time.Time,
		to time.Time,
	) ([]schema.MorningCall, error)
	ReadEveningCalls(
		userIdStr schema.CleanText,
		from time.Time,
		to time.Time,
	) ([]schema.EveningCall, error)
	ReadDailyReport(repoIdStr schema.CleanText) (*schema.DailyReport, error)
	ReadDailyReports(
		userIdStr schema.CleanText,
		from time.Time,
		to time.Time,
	) ([]schema.DailyReport, error)
	ReadLocations(repoIdStr schema.CleanText) ([]schema.Location, error)
	ReadRestings(repoIdStr schema.CleanText) ([]schema.Resting, error)
	ReadWaitings(repoIdStr schema.CleanText) ([]schema.Waiting, error)
	ReadLoadings(repoIdStr schema.CleanText) ([]schema.Loading, error)

	Tx(txf func(Transaction) error) error
	Once() Transaction
}
