package Date

import "time"

// StartOfYear 获取日期中当年的开始时间.
func (*Date)StartOfYear(date time.Time) time.Time {
	return time.Date(date.Year(), 1, 1, 0, 0, 0, 0, date.Location())
}
