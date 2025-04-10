package q

import (
	"nippo/schema"
	"nippo/view/viewprops"
)

type QPrint struct {
	repoIdStr schema.CleanText
}

func (q QPrint) Query(props *viewprops.Props) ([]byte, error) {
	db := props.Db()

	report, err := db.ReadDailyReport(q.repoIdStr)
}
