package schema

import "math"

type StateIdx uint32

const ST_USR_NOT_CONFIGURED StateIdx = math.MaxUint32 //4294967295

const (
	ST_IDX_NONE StateIdx = iota
)
