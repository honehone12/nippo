package q

import (
	"nippo/schema"
	"nippo/view/viewprops"
	"time"
)

type QCalls struct {
	userIdStr schema.CleanText
	year      int
	month     time.Month
}

func (q QCalls) Query(props *viewprops.Props) ([]byte, error) {
	panic("not implemented")
}
