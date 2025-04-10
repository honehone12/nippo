package q

import (
	"encoding/json"
	"errors"
	"nippo/schema"
	"nippo/view/viewprops"

	"gorm.io/gorm"
)

type QPrint struct {
	repoIdStr schema.CleanText
}

func (q QPrint) Query(props *viewprops.Props) ([]byte, error) {
	db := props.Db()

	report, err := db.ReadDailyReport(q.repoIdStr)
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

	var mCall *schema.MorningCall
	if len(report.MorningCallId) > 0 {
		mCall, err = db.ReadMorningCall(report.MorningCallId)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	var eCall *schema.EveningCall
	if len(report.EveningCallId) > 0 {
		eCall, err = db.ReadEveningCall(report.EveningCallId)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	locs, err := db.ReadLocations(q.repoIdStr)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		locs = []schema.Location{}
	} else if err != nil {
		return nil, err
	}

	waits, err := db.ReadWaitings(q.repoIdStr)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		waits = []schema.Waiting{}
	} else if err != nil {
		return nil, err
	}

	loads, err := db.ReadLoadings(q.repoIdStr)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		loads = []schema.Loading{}
	} else if err != nil {
		return nil, err
	}

	rests, err := db.ReadRestings(q.repoIdStr)
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
