package timeconv

import (
	"testing"
	"time"
)

func TestUtcRangeFromJstMonth(t *testing.T) {
	from, to := UtcRangeFromJstMonth(2025, time.April)

	if from.Year() != 2025 || from.Month() != time.March || from.Day() != 31 || from.Hour() != 15 {
		t.Fatalf("wrong time")
	}
	if to.Year() != 2025 || to.Month() != time.April || to.Day() != 30 || to.Hour() != 15 {
		t.Fatalf("wrong time")
	}
}
