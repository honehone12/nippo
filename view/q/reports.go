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

type QReports struct {
	userIdStr schema.CleanText
	year      int
	month     time.Month
}

func (q QReports) Query(props *viewprops.Props) ([]byte, error) {
	from, to := timeconv.UtcRangeFromJstMonth(q.year, q.month)

	reports, err := props.Db().ReadDailyReports(q.userIdStr, from, to)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		reports = []schema.DailyReport{}
	} else if err != nil {
		return nil, err
	}

	return json.Marshal(reports)
}
