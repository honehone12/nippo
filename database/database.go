package database

import (
	"nippo/schema"
)

type Database interface {
	CreateOrg(name schema.CleanText) ([]uint8, error)
	ReadOrg(idStr string) (*schema.Org, error)
	UpdateOrgLineId(idStr string, newLineId schema.CleanText) error

	CreateUser(orgIdStr string, lineId schema.CleanText) ([]uint8, error)
	ReadUser(lineId schema.CleanText) (*schema.User, error)
	UpdateUserName(id []uint8, name schema.CleanText) error
	UpdateUserCarNumber(id []uint8, carNumber schema.CleanText) error
	UpdateUserCurrentOpe(id []uint8, call []uint8) error
	UpdateUserCurrentDuty(id []uint8, duty []uint8) error
	UpdateUserNextState(id []uint8, nextState schema.StateIdx) error

	CreateMorningCall(userId []uint8, photo string) ([]uint8, error)
	UpdateMoCallHealth(callId []uint8) error
	UpdateMoCallCarCheck(CallId []uint8) error
	UpdateMoCallNote(callId []uint8, note schema.CleanText) error

	CreateEveningCall(userId []uint8, photo string) ([]uint8, error)
	UpdateEvCallNote(callId []uint8, note schema.CleanText) error

	CreateDailyReport(userId []uint8, morningId []uint8) ([]uint8, error)
	UpdateDReportEvCall(repoId []uint8, evId []uint8) error
	UpdateDReportMtrPhoto(repoId []uint8, photo string) error
	UpdateDReportTrouble(repoId []uint8, trouble schema.CleanText) error
	UpdateDReportNote(repoId []uint8, note schema.CleanText) error
}
