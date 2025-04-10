package q

import (
	"encoding/json"
	"errors"
	"nippo/schema"
	"nippo/timeconv"
	"nippo/view/viewprops"
	"time"

	"gorm.io/gorm"
)

type QCalls struct {
	userIdStr schema.CleanText
	year      int
	month     time.Month
}

func (q QCalls) Query(props *viewprops.Props) ([]byte, error) {
	db := props.Db()
	from, to := timeconv.UtcRangeFromJstMonth(q.year, q.month)

	mCalls, err := db.ReadMorningCalls(q.userIdStr, from, to)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		mCalls = []schema.MorningCall{}
	} else if err != nil {
		return nil, err
	}

	eCalls, err := db.ReadEveningCalls(q.userIdStr, from, to)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		eCalls = []schema.EveningCall{}
	} else if err != nil {
		return nil, err
	}

	return json.Marshal(schema.Calls{
		MorningCalls: mCalls,
		EveningCalls: eCalls,
	})
}
