package q

import (
	"nippo/schema"
	"nippo/view/viewprops"
	"time"
)

type QReports struct {
	userIdStr schema.CleanText
	year      int
	month     time.Month
}

func (q QReports) Query(props *viewprops.Props) ([]byte, error) {
	panic("not implemented")
}
