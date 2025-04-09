package database

import "nippo/schema"

type Database interface {
	ReadOrg(idStr string) (*schema.Org, error)
	ReadUser(lineId schema.CleanText) (*schema.User, error)
	ReadUsers(orgIdStr string) ([]schema.User, error)

	Tx(txf func(Transaction) error) error
	Once() Transaction
}
