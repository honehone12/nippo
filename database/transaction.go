package database

import "nippo/schema"

type Transaction interface {
	CreateOrg(name schema.CleanText) ([]uint8, error)
	UpdateOrgLineId(idStr string, newLineId schema.CleanText) error

	CreateUser(orgIdStr string, lineId schema.CleanText) ([]uint8, error)
	UpdateUserName(userId []uint8, name schema.CleanText) error
	UpdateUserCarNumber(userId []uint8, carNumber schema.CleanText) error
	UpdateUserCurrentOpe(userId []uint8, opeId []uint8) error
	UpdateUserCurrentDuty(userId []uint8, repoId []uint8) error
	UpdateUserNextState(userId []uint8, nextState schema.StateIdx) error

	CreateMorningCall(userId []uint8, photo string) ([]uint8, error)
	UpdateMCallHealth(callId []uint8) error
	UpdateMCallCarCheck(CallId []uint8) error
	UpdateMCallNote(callId []uint8, note schema.CleanText) error

	CreateEveningCall(userId []uint8, photo string) ([]uint8, error)
	UpdateECallNote(callId []uint8, note schema.CleanText) error

	CreateDailyReport(userId []uint8, morningId []uint8) ([]uint8, error)
	UpdateDReportEvCall(repoId []uint8, eveningId []uint8) error
	UpdateDReportMtrPhoto(repoId []uint8, photo string) error
	UpdateDReportTrouble(repoId []uint8, trouble schema.CleanText) error
	UpdateDReportNote(repoId []uint8, note schema.CleanText) error

	CreateLocation(
		repoId []uint8,
		label schema.CleanText,
		address schema.CleanText,
		latitude float64,
		longitude float64,
	) ([]uint8, error)
	UpdateLocNote(locId []uint8, shortNote schema.CleanText) error

	CreateResting(
		repoId []uint8,
		label schema.CleanText,
		address schema.CleanText,
		latitude float64,
		longitude float64,
	) ([]uint8, error)
	UpdateRestNote(restId []uint8, shortNote schema.CleanText) error
}
