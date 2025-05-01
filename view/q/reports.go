package q

import (
	"bytes"
	"encoding/json"
	"errors"
	"nippo/schema"
	"nippo/timeconv"
	"nippo/view/viewprops"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QReports struct {
	userId uuid.UUID
	year   int
	month  time.Month
}

func (q QReports) Query(props *viewprops.Props) ([]byte, error) {
	db := props.Db()
	from, to := timeconv.UtcRangeFromJstMonth(q.year, q.month)
	userId := q.userId[:]

	user, err := db.ReadUser(userId)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(user.OrgId, props.OrgId()) {
		return nil, errors.New("viewing wrong org")
	}

	reports, err := db.ReadDailyReports(userId, from, to)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		reports = []schema.DailyReport{}
	} else if err != nil {
		return nil, err
	}

	return json.Marshal(reports)
}
