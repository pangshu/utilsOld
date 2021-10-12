package Date

import "time"

// MonthDays 获取指定月份的天数.years年份,可选,默认当前年份.
func (*Date)MonthDays(month int, years ...int) int {
    if month < 1 || month > 12 {
        return 0
    }

    var year int
    var days int
    if len(years) == 0 {
        year = time.Now().Year()
    } else {
        year = years[0]
    }

    if month != 2 {
        if month == 4 || month == 6 || month == 9 || month == 11 {
            days = 30
        } else {
            days = 31
        }
    } else {
        if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
            days = 29
        } else {
            days = 28
        }
    }

    return days
}