package time

import (
	"time"
)

// InOpenDayRange 判断当前时间是否在 start 和 end 时间段内
//
// start 和 end 格式目前只支持 hh:mm
func InOpenDayRange(nowUnix int64, startStr, endStr string) bool {
	now := time.Unix(nowUnix, 0)
	const layout = "15:04"
	start, err := time.Parse(layout, startStr)
	if err != nil {
		return false
	}
	end, err := time.Parse(layout, endStr)
	if err != nil {
		return false
	}
	now = time.Date(start.Year(), start.Month(), start.Day(), now.Hour(), now.Minute(), 0, 0, time.UTC)
	afterStart := now.Equal(start) || now.After(start)
	beforeEnd := now.Equal(end) || now.Before(end)
	return afterStart && beforeEnd
}
