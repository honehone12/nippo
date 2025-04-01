package state

import "nippo/schema"

const (
	ST_USER_NOT_CONFIGURED schema.StateIdx = iota

	ST_USER_SET_NAME
	ST_USER_SET_CAR

	ST_IDX_NONE
)
