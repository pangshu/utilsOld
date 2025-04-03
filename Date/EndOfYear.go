package Date

import "time"

// EndOfYear 获取日期中当年的结束时间.
func (toolDate *Date)EndOfYear(date time.Time) time.Time {
	return time.Date(date.Year(), 1, 1, 0, 0, 0, 0, date.Location()).AddDate(1, 0, 0).Add(-time.Nanosecond)
}