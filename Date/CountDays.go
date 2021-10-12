package Date

import "time"

// CountDays 获取两个日期的间隔天数.
func (*Date)CountDays(fromDate, toDate time.Time) int {
    return int(toDate.Sub(fromDate) / (24 * time.Hour))
}