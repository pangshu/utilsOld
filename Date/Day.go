package Date

import "time"

// Day 获取日份.
func (*Date)Day(t ...time.Time) int {
    var tm time.Time
    if len(t) > 0 {
        tm = t[0]
    } else {
        tm = time.Now()
    }
    return tm.Day()
}
