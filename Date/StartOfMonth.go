package Date

import "time"

// StartOfMonth 获取日期中当月的开始时间.
func (*Date)StartOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}
