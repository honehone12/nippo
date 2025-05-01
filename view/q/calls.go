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

type QCalls struct {
	userId uuid.UUID
	year   int
	month  time.Month
}

func (q QCalls) Query(props *viewprops.Props) ([]byte, error) {
	db := props.Db()
	userId := q.userId[:]

	user, err := db.ReadUser(userId)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(user.OrgId, props.OrgId()) {
		return nil, errors.New("viewing wrong org")
	}

	from, to := timeconv.UtcRangeFromJstMonth(q.year, q.month)

	mCalls, err := db.ReadMorningCalls(userId, from, to)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		mCalls = []schema.MorningCall{}
	} else if err != nil {
		return nil, err
	}

	eCalls, err := db.ReadEveningCalls(userId, from, to)
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
