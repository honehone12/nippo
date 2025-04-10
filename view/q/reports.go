package q

import (
	"encoding/json"
	"nippo/schema"
	"nippo/timeconv"
	"nippo/view/viewprops"
	"time"
)

type QReports struct {
	userIdStr schema.CleanText
	year      int
	month     time.Month
}

func (q QReports) Query(props *viewprops.Props) ([]byte, error) {
	from, to := timeconv.UtcRangeFromJstMonth(q.year, q.month)

	reports, err := props.Db().ReadDailyReports(q.userIdStr, from, to)
	if err != nil {
		return nil, err
	}

	return json.Marshal(reports)
}
