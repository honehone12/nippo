package database

import (
	"nippo/schema"
	"time"
)

type Database interface {
	ReadOrg(id []uint8) (*schema.Org, error)
	ReadOrgLineId(id []uint8) (schema.CleanText, error)
	ReadOrgEmail(id []uint8) (schema.CleanText, error)

	ReadUserEmailOrgId(id []uint8) (*schema.User, error)
	ReadUser_Line(lineId schema.CleanText) (*schema.User, error)
	ReadUserOrgId(id []uint8) ([]uint8, error)
	ReadUserOrgId_Email(email schema.CleanText) ([]uint8, error)
	ReadUserBasic_Email(email schema.CleanText) (*schema.User, error)
	ReadUserEmail(id []uint8) (schema.CleanText, error)
	ReadUsersBasic(orgId []uint8) ([]schema.User, error)

	ReadMorningCall(id []uint8) (*schema.MorningCall, error)
	ReadMCallCarNum(id []uint8) (schema.CleanText, error)
	ReadMCallAlcPhoto(id []uint8) (string, error)
	ReadMorningCalls(
		userId []uint8,
		from time.Time,
		to time.Time,
	) ([]schema.MorningCall, error)

	ReadEveningCall(id []uint8) (*schema.EveningCall, error)
	ReadECallAlcPhoto(id []uint8) (string, error)
	ReadEveningCalls(
		userId []uint8,
		from time.Time,
		to time.Time,
	) ([]schema.EveningCall, error)

	ReadDailyReport(id []uint8) (*schema.DailyReport, error)
	ReadDRepoDownload(id []uint8) (*schema.DailyReport, error)
	ReadDRepoCarNum(id []uint8) (schema.CleanText, error)
	ReadDailyReports(
		userId []uint8,
		from time.Time,
		to time.Time,
	) ([]schema.DailyReport, error)

	ReadSites(repoId []uint8) ([]schema.Site, error)
	ReadLocations(repoId []uint8) ([]schema.Location, error)
	ReadRestings(repoId []uint8) ([]schema.Resting, error)
	ReadWaitings(repoId []uint8) ([]schema.Waiting, error)
	ReadLoadings(repoId []uint8) ([]schema.Loading, error)

	Tx(txf func(Transaction) error) error
	Once() Transaction
}
