package state

type State interface {
	run() (StateIdx, error)
}
