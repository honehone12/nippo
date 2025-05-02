package q

import (
	"bytes"
	"encoding/json"
	"errors"
	"nippo/schema"
	"nippo/view/viewprops"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QPrint struct {
	repoId uuid.UUID
}

func (q QPrint) Query(props *viewprops.Props, email schema.CleanText) ([]byte, error) {
	db := props.Db()
	repoId := q.repoId[:]

	report, err := db.ReadDailyReport(repoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return json.Marshal(schema.DailyReportPrint{
			Locations: []schema.Location{},
			Waitings:  []schema.Waiting{},
			Loadings:  []schema.Loading{},
			Resings:   []schema.Resting{},
		})
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

	var mCall *schema.MorningCall
	if len(report.MorningCallId) > 0 {
		mCall, err = db.ReadMorningCall(report.MorningCallId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	var eCall *schema.EveningCall
	if len(report.EveningCallId) > 0 {
		eCall, err = db.ReadEveningCall(report.EveningCallId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	locs, err := db.ReadLocations(repoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		locs = []schema.Location{}
	} else if err != nil {
		return nil, err
	}

	waits, err := db.ReadWaitings(repoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		waits = []schema.Waiting{}
	} else if err != nil {
		return nil, err
	}

	loads, err := db.ReadLoadings(repoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		loads = []schema.Loading{}
	} else if err != nil {
		return nil, err
	}

	rests, err := db.ReadRestings(repoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		rests = []schema.Resting{}
	} else if err != nil {
		return nil, err
	}

	return json.Marshal(schema.DailyReportPrint{
		DailyReport: report,
		MorningCall: mCall,
		EveningCall: eCall,
		Locations:   locs,
		Waitings:    waits,
		Loadings:    loads,
		Resings:     rests,
	})
}
