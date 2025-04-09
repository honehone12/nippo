package q

import (
	"nippo/view/viewprops"
	"time"
)

type QReports struct {
	year  int
	month time.Month
}

func (q QReports) Query(props *viewprops.Props) ([]byte, error) {
	panic("not implemented")
}
