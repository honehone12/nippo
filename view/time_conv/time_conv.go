package timeconv

import "time"

var jST = time.FixedZone("JST", 9*60*60)

func UtcRangeFromJstMonth(year int, jstMonth time.Month) (time.Time, time.Time) {
	jstFrom := time.Date(year, jstMonth, 1, 0, 0, 0, 0, jST)
	jstTo := jstFrom.AddDate(0, 1, 0)

	return jstFrom.UTC(), jstTo.UTC()
}
