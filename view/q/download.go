package q

import (
	"encoding/json"
	"errors"
	"nippo/schema"
	"nippo/view/viewprops"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QDownload struct {
	repoId uuid.UUID
}

func (q QDownload) Query(props *viewprops.Props) ([]byte, error) {
	db := props.Db()
	repoId := q.repoId[:]

	report, err := db.ReadDailyReport(repoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return json.Marshal(schema.Photos{})
	} else if err != nil {
		return nil, err
	}

	var meter string
	if len(report.MeterPhoto) > 0 {
		meter = report.MeterPhoto
	}

	var morningAlc string
	if len(report.MorningCallId) > 0 {
		mCall, err := db.ReadMorningCall(report.MorningCallId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else if err == nil && len(mCall.AlcPhoto) > 0 {
			morningAlc = mCall.AlcPhoto
		}
	}

	var eveningAlc string
	if len(report.EveningCallId) > 0 {
		eCall, err := db.ReadEveningCall(report.EveningCallId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else if err == nil && len(eCall.AlcPhoto) > 0 {
			eveningAlc = eCall.AlcPhoto
		}
	}

	return json.Marshal(schema.Photos{
		MorningAlc: morningAlc,
		EveningAlc: eveningAlc,
		Meter:      meter,
	})
}
