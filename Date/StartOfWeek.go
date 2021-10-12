package Date

import "time"

// StartOfWeek 获取日期中当周的开始时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func (*Date)StartOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	weekStart := time.Monday
	if len(weekStartDay) > 0 {
		weekStart = weekStartDay[0]
	}

	// 当前是周几
	weekday := int(date.Weekday())
	if weekStart != time.Sunday {
		weekStartDayInt := int(weekStart)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}

	return time.Date(date.Year(), date.Month(), date.Day()-weekday, 0, 0, 0, 0, date.Location())
}

