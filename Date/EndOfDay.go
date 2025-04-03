package Date

import "time"

// EndOfDay 获取日期中当天的结束时间.
func (*Date)EndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), date.Location())
}

