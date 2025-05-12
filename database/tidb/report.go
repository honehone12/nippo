package tidb

import (
	"errors"
	"nippo/schema"
	"time"

	"github.com/google/uuid"
)

func (ti TiDbTx) CreateDailyReport(
	userId []uint8,
	name schema.CleanText,
	carNum schema.CleanText,
	morningId []uint8,
) ([]uint8, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}
	if len(morningId) != schema.ID_LEN {
		return nil, errors.New("unexpected morning id len")
	}
	if len(name) <= 1 || len(name) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected name length")
	}
	if len(carNum) == 0 || len(carNum) > schema.MAX_SHORT_TEXT_LEN {
		return nil, errors.New("unexpected car number length")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.DailyReport{
		Id:            id[:],
		MorningCallId: morningId,
		Name:          name,
		CarNumber:     carNum,
	}
	user := schema.User{Id: userId}
	err = ti.orm.Model(&user).Association("DailyReports").Append(&call)
	if err != nil {
		return nil, err
	}

	return call.Id, nil
}

func (ti TiDb) ReadDailyReport(id []uint8) (*schema.DailyReport, error) {
	if len(id) != schema.ID_LEN {
		return nil, errors.New("unexpected report id length")
	}

	repo := schema.DailyReport{Id: id}
	if res := ti.orm.Take(&repo); res.Error != nil {
		return nil, res.Error
	}

	return &repo, nil
}

func (ti TiDb) ReadDRepoCarNum(id []uint8) (schema.CleanText, error) {
	if len(id) != schema.ID_LEN {
		return "", errors.New("unexpected report id length")
	}

	repo := schema.DailyReport{Id: id}
	if res := ti.orm.Select("car_number").Take(&repo); res.Error != nil {
		return "", res.Error
	}

	return repo.CarNumber, nil
}

func (ti TiDb) ReadDailyReports(
	userId []uint8,
	from time.Time,
	to time.Time,
) ([]schema.DailyReport, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}

	user := schema.User{Id: userId}
	var reports []schema.DailyReport
	err := ti.orm.Model(&user).Where(
		"created_at >= ? AND created_at < ?",
		from,
		to,
	).Association("DailyReports").Find(&reports)
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (ti TiDbTx) UpdateDReportECall(repoId []uint8, eveningId []uint8) error {
	if len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}
	if len(eveningId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}

	repo := schema.DailyReport{Id: repoId}
	res := ti.orm.Model(&repo).Update("evening_call_id", eveningId)
	return res.Error
}

func (ti TiDbTx) UpdateDReportMMtr(repoId []uint8, meter uint64) error {
	if len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}

	repo := schema.DailyReport{Id: repoId}
	res := ti.orm.Model(&repo).Update("morning_meter", meter)
	return res.Error
}

func (ti TiDbTx) UpdateDReportMMtrPhoto(repoId []uint8, photo string) error {
	if len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}
	if len(photo) > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected photo length")
	}

	repo := schema.DailyReport{Id: repoId}
	res := ti.orm.Model(&repo).Update("morning_meter_photo", photo)
	return res.Error
}

func (ti TiDbTx) UpdateDReportEMtr(repoId []uint8, meter uint64) error {
	if len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}

	repo := schema.DailyReport{Id: repoId}
	res := ti.orm.Model(&repo).Update("evening_meter", meter)
	return res.Error
}

func (ti TiDbTx) UpdateDReportEMtrPhoto(repoId []uint8, photo string) error {
	if len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}
	if len(photo) > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected photo length")
	}

	repo := schema.DailyReport{Id: repoId}
	res := ti.orm.Model(&repo).Update("evening_meter_photo", photo)
	return res.Error
}

func (ti TiDbTx) UpdateDReportTrouble(repoId []uint8, trouble schema.CleanText) error {
	if len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}
	if len(trouble) > schema.MAX_LONG_TEXT_LEN {
		return errors.New("unexpected trouble length")
	}

	repo := schema.DailyReport{Id: repoId}
	res := ti.orm.Model(&repo).Update("trouble", trouble)
	return res.Error
}

func (ti TiDbTx) UpdateDReportNote(repoId []uint8, note schema.CleanText) error {
	if len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}
	if len(note) > schema.MAX_LONG_TEXT_LEN {
		return errors.New("unexpected note length")
	}

	repo := schema.DailyReport{Id: repoId}
	res := ti.orm.Model(&repo).Update("note", note)
	return res.Error
}
