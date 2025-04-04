package schema

type CallMethod uint32

const (
	METHOD_PERSON CallMethod = iota + 1
	METHOD_REMOTE
)

type StateIdx uint32

type CleanText string
