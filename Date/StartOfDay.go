package Date

import "time"

// StartOfDay 获取日期中当天的开始时间.
func (*Date)StartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}
