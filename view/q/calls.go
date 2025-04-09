package q

import (
	"nippo/view/viewprops"
	"time"
)

type QCalls struct {
	year  int
	month time.Month
}

func (q QCalls) Query(props *viewprops.Props) ([]byte, error) {
	panic("not implemented")
}
