package time

import (
	"testing"
	"time"
)

func TestWeekDay(t *testing.T) {
	now := time.Now()
	/*** 需注意 time.WeekDay 的周日是 0 ***/
	t.Logf("now: %v", now.Weekday())
}
