package database

import (
	"nippo/schema"
)

type Database interface {
	CreateOrg(name string) ([]uint8, error)
	ReadOrg(idStr string) (*schema.Org, error)
	UpdateOrgLineId(idStr string, newLineId string) error

	CreateUser(
		orgIdStr string,
		name string,
		lineId string,
	) ([]uint8, error)
	ReadUser(lineId string) (*schema.User, error)
	UpdateUserCarNumber(lineId string, carNumber string) error
	UpdateUserNextState(lineId string, nextState schema.StateIdx) error
	UpdateUserCurrentDuty(lineId string, duty []uint8) error

	CreateMorningCall(userLineId string) ([]uint8, error)

	CreateDailyReport(userLineId string, morningId []uint8) ([]uint8, error)

	CreateEveningCall(userLineId string) ([]uint8, error)
}
