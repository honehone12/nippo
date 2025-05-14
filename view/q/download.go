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

func (q QDownload) Query(props *viewprops.Props, viewer *viewprops.Viewer) ([]byte, error) {
	db := props.Db()
	repoId := q.repoId[:]
	photos := schema.Photos{
		MorningAlc: []byte{},
		EveningAlc: []byte{},
		MorningMtr: []byte{},
		EveningMtr: []byte{},
	}

	report, err := db.ReadDRepoDownload(repoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return json.Marshal(photos)
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

	if len(report.MorningMeterPhoto) > 0 {
		url, err := sss.GetPresigned(props, report.MorningMeterPhoto)
		if err != nil {
			return nil, err
		}

		photos.MorningMtr = []byte(url)
	}

	if len(report.EveningMeterPhoto) > 0 {
		url, err := sss.GetPresigned(props, report.EveningMeterPhoto)
		if err != nil {
			return nil, err
		}

		photos.EveningMtr = []byte(url)
	}

	if len(report.MorningCallId) > 0 {
		alc, err := db.ReadMCallAlcPhoto(report.MorningCallId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else if len(alc) > 0 {
			url, err := sss.GetPresigned(props, alc)
			if err != nil {
				return nil, err
			}

			photos.MorningAlc = []byte(url)
		}
	}

	if len(report.EveningCallId) > 0 {
		alc, err := db.ReadECallAlcPhoto(report.EveningCallId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else if len(alc) > 0 {
			url, err := sss.GetPresigned(props, alc)
			if err != nil {
				return nil, err
			}

			photos.EveningAlc = []byte(url)
		}
	}

	return json.Marshal(photos)
}
