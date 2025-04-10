package q

import (
	"encoding/json"
	"nippo/schema"
	"nippo/view/viewprops"
)

type QPrint struct {
	repoIdStr schema.CleanText
}

func (q QPrint) Query(props *viewprops.Props) ([]byte, error) {
	db := props.Db()

	report, err := db.ReadDailyReport(q.repoIdStr)
	if err != nil {
		return nil, err
	}

	locs, err := db.ReadLocations(q.repoIdStr)
	if err != nil {
		return nil, err
	}

	waits, err := db.ReadWaitings(q.repoIdStr)
	if err != nil {
		return nil, err
	}

	loads, err := db.ReadLoadings(q.repoIdStr)
	if err != nil {
		return nil, err
	}

	rests, err := db.ReadRestings(q.repoIdStr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(schema.DailyReportPrint{
		DailyReport: *report,
		MorningCall: schema.MorningCall{},
		EveningCall: schema.EveningCall{},
		Locations:   locs,
		Waitings:    waits,
		Loadings:    loads,
		Resings:     rests,
	})
}
