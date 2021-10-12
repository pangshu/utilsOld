package Date

import "time"

// EndOfMonth 获取日期中当月的结束时间.
func (toolDate *Date)EndOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location()).AddDate(0, 1, 0).Add(-time.Nanosecond)
}
