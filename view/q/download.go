package q

import (
	"bytes"
	"encoding/json"
	"errors"
	"nippo/schema"
	"nippo/sss"
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

	orgId, err := db.ReadUserOrgId(report.UserId)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(orgId, props.OrgId()) {
		return nil, errors.New("viewing wrong org")
	}

	var meter []byte
	if len(report.MeterPhoto) > 0 {
		url, err := sss.GetPresigned(props, report.MeterPhoto)
		if err != nil {
			return nil, err
		}

		meter = []byte(url)
	}

	var morningAlc []byte
	if len(report.MorningCallId) > 0 {
		mCall, err := db.ReadMorningCall(report.MorningCallId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else if err == nil && len(mCall.AlcPhoto) > 0 {
			url, err := sss.GetPresigned(props, mCall.AlcPhoto)
			if err != nil {
				return nil, err
			}

			morningAlc = []byte(url)
		}
	}

	var eveningAlc []byte
	if len(report.EveningCallId) > 0 {
		eCall, err := db.ReadEveningCall(report.EveningCallId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else if err == nil && len(eCall.AlcPhoto) > 0 {
			url, err := sss.GetPresigned(props, eCall.AlcPhoto)
			if err != nil {
				return nil, err
			}

			eveningAlc = []byte(url)
		}
	}

	return json.Marshal(schema.Photos{
		MorningAlc: morningAlc,
		EveningAlc: eveningAlc,
		Meter:      meter,
	})
}
