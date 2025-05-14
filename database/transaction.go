package database

import (
	"nippo/schema"
	"time"
)

type Transaction interface {
	CreateOrg(name schema.CleanText, email schema.CleanText) ([]uint8, error)
	UpdateOrgLineId(id []uint8, newLineId schema.CleanText) error

	CreateUser(orgId []uint8, lineId schema.CleanText) ([]uint8, error)
	UpdateUserName(userId []uint8, name schema.CleanText) error
	UpdateUserSubAdmin(userId []uint8, subAdmin bool) error
	UpdateUserEmail(userId []uint8, email schema.CleanText) error
	UpdateUserCurrentOpe(userId []uint8, opeId []uint8) error
	UpdateUserCurrentDuty(userId []uint8, repoId []uint8) error
	UpdateUserNextState(userId []uint8, nextState schema.StateIdx) error

	CreateMorningCall(
		userId []uint8,
		name schema.CleanText,
		photo string,
	) ([]uint8, error)
	UpdateMCallHealth(callId []uint8) error
	UpdateMCallCarCheck(CallId []uint8) error
	UpdateMCallCarNumber(callId []uint8, carNumber schema.CleanText) error
	UpdateMCallLicenseExp(callId []uint8, exp time.Time) error
	UpdateMCallNote(callId []uint8, note schema.CleanText) error

	CreateEveningCall(
		userId []uint8,
		name schema.CleanText,
		carNum schema.CleanText,
		photo string,
	) ([]uint8, error)
	UpdateECallNote(callId []uint8, note schema.CleanText) error

	CreateDailyReport(
		userId []uint8,
		name schema.CleanText,
		carNum schema.CleanText,
		morningId []uint8,
	) ([]uint8, error)
	UpdateDReportECall(repoId []uint8, eveningId []uint8) error
	UpdateDReportMMtr(repoId []uint8, meter uint64) error
	UpdateDReportMMtrPhoto(repoId []uint8, photo string) error
	UpdateDReportEMtr(repoId []uint8, meter uint64) error
	UpdateDReportEMtrPhoto(repoId []uint8, photo string) error
	UpdateDReportNonDutyDst(repoId []uint8, dst uint64) error
	UpdateDReportTrouble(repoId []uint8, trouble schema.CleanText) error
	UpdateDReportNote(repoId []uint8, note schema.CleanText) error

	CreateSite(repoId []uint8, name schema.CleanText) ([]uint8, error)
	UpdateSiteNote(siteId []uint8, shortNote schema.CleanText) error

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

	CreateWaiting(
		repoId []uint8,
		label schema.CleanText,
		address schema.CleanText,
		latitude float64,
		longitude float64,
	) ([]uint8, error)
	UpdateWaitNote(waitId []uint8, shortNote schema.CleanText) error

	CreateLoading(
		repoId []uint8,
		label schema.CleanText,
		address schema.CleanText,
		latitude float64,
		longitude float64,
	) ([]uint8, error)
	UpdateLoadShipCheck(loadId []uint8, check bool) error
	UpdateLoadNote(loadId []uint8, shortNote schema.CleanText) error
}
