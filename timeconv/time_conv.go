package timeconv

import (
	"errors"
	"strconv"
	"time"
)

var jST = time.FixedZone("JST", 9*60*60)

func AToYear(y string) (int, error) {
	year, err := strconv.Atoi(y)
	if err != nil {
		return 0, err
	}
	if year < 2025 || year > 2099 {
		return 0, errors.New("out of range")
	}

	return year, nil
}

func AToMonth(m string) (time.Month, error) {
	month, err := strconv.Atoi(m)
	if err != nil {
		return time.Month(0), err
	}
	if month < 1 || month < 12 {
		return time.Month(0), errors.New("out of range")
	}

	return time.Month(month), nil
}

func AToYM(y string, m string) (int, time.Month, error) {
	year, err := AToYear(y)
	if err != nil {
		return 0, time.Month(0), err
	}

	month, err := AToMonth(m)
	if err != nil {
		return 0, time.Month(0), err
	}

	return year, month, nil
}

func UtcRangeFromJstMonth(year int, jstMonth time.Month) (time.Time, time.Time) {
	jstFrom := time.Date(year, jstMonth, 1, 0, 0, 0, 0, jST)
	jstTo := jstFrom.AddDate(0, 1, 0)

	return jstFrom.UTC(), jstTo.UTC()
}
