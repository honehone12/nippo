package tidb

import (
	"errors"
	"nippo/schema"

	"github.com/google/uuid"
)

func (ti *TiDbTx) CreateDailyReport(userId []uint8, morningId []uint8) ([]uint8, error) {
	if len(userId) != schema.ID_LEN {
		return nil, errors.New("unexpected user id length")
	}
	if len(morningId) != schema.ID_LEN {
		return nil, errors.New("unexpected morning id len")
	}

	user := schema.User{}
	if res := ti.orm.Select(
		"id",
		"name",
		"car_number",
	).Take(&user, "id = ?", userId); res.Error != nil {
		return nil, res.Error
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	call := schema.DailyReport{
		Id:            id[:],
		MorningCallId: morningId,
		Name:          user.Name,
		CarNumber:     user.CarNumber,
	}
	if err := ti.orm.Model(&user).Association("DailyReports").Append(&call); err != nil {
		return nil, err
	}

	return call.Id, nil
}

func (ti *TiDbTx) UpdateDReportEvCall(repoId []uint8, evId []uint8) error {
	if len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}
	if len(evId) != schema.ID_LEN {
		return errors.New("unexpected call id length")
	}

	repo := schema.DailyReport{}
	res := ti.orm.Model(&repo).Where("id = ?", repoId).Update("evening_call_id", evId)
	return res.Error
}

func (ti *TiDbTx) UpdateDReportMtrPhoto(repoId []uint8, photo string) error {
	if len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}
	if len(photo) > schema.MAX_SHORT_TEXT_LEN {
		return errors.New("unexpected photo length")
	}

	repo := schema.DailyReport{}
	res := ti.orm.Model(&repo).Where("id = ?", repoId).Update("meter_photo", photo)
	return res.Error
}

func (ti *TiDbTx) UpdateDReportTrouble(repoId []uint8, trouble schema.CleanText) error {
	if len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}
	if len(trouble) > schema.MAX_LONG_TEXT_LEN {
		return errors.New("unexpected trouble length")
	}

	repo := schema.DailyReport{}
	res := ti.orm.Model(&repo).Where("id = ?", repoId).Update("trouble", trouble)
	return res.Error
}

func (ti *TiDbTx) UpdateDReportNote(repoId []uint8, note schema.CleanText) error {
	if len(repoId) != schema.ID_LEN {
		return errors.New("unexpected report id length")
	}
	if len(note) > schema.MAX_LONG_TEXT_LEN {
		return errors.New("unexpected note length")
	}

	repo := schema.DailyReport{}
	res := ti.orm.Model(&repo).Where("id = ?", repoId).Update("note", note)
	return res.Error
}
