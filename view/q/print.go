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

func (q QPrint) Query(props *viewprops.Props, viewer *viewprops.Viewer) ([]byte, error) {
	db := props.Db()
	repoId := q.repoId[:]
	reports := schema.DailyReportPrint{
		Sites:     []schema.Site{},
		Locations: []schema.Location{},
		Waitings:  []schema.Waiting{},
		Loadings:  []schema.Loading{},
		Resings:   []schema.Resting{},
	}

	report, err := db.ReadDailyReport(repoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return json.Marshal(reports)
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

	if !viewer.Privileged() && !bytes.Equal(report.UserId, viewer.Id()) {
		return nil, errors.New("no permission for viewing other user")
	}

	reports.DailyReport = report

	if len(report.MorningCallId) > 0 {
		mCall, err := db.ReadMorningCall(report.MorningCallId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else if mCall != nil {
			reports.MorningCall = mCall
		}
	}

	if len(report.EveningCallId) > 0 {
		eCall, err := db.ReadEveningCall(report.EveningCallId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else if eCall != nil {
			reports.EveningCall = eCall
		}
	}

	sites, err := db.ReadSites(repoId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if len(sites) > 0 {
		reports.Sites = append(reports.Sites, sites...)
	}

	locs, err := db.ReadLocations(repoId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if len(locs) > 0 {
		reports.Locations = append(reports.Locations, locs...)
	}

	waits, err := db.ReadWaitings(repoId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if len(waits) > 0 {
		reports.Waitings = append(reports.Waitings, waits...)
	}

	loads, err := db.ReadLoadings(repoId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if len(loads) > 0 {
		reports.Loadings = append(reports.Loadings, loads...)
	}

	rests, err := db.ReadRestings(repoId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if len(rests) > 0 {
		reports.Resings = append(reports.Resings, rests...)
	}

	return json.Marshal(reports)
}
