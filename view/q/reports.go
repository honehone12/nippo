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

func (q QReports) Query(props *viewprops.Props, viewer *viewprops.Viewer) ([]byte, error) {
	db := props.Db()
	userId := q.userId[:]

	orgId, err := db.ReadUserOrgId(userId)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(orgId, props.OrgId()) {
		return nil, errors.New("viewing wrong org")
	}

	if !viewer.Privileged() && !bytes.Equal(userId, viewer.Id()) {
		return nil, errors.New("no permission for viewing other user")
	}

	from, to := timeconv.UtcRangeFromJstMonth(q.year, q.month)

	reports, err := db.ReadDailyReports(userId, from, to)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		reports = []schema.DailyReport{}
	} else if err != nil {
		return nil, err
	}

	return json.Marshal(reports)
}
