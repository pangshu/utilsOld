package Date

import "time"

// Month 获取月份.
func (*Date)Month(t ...time.Time) int {
    var tm time.Time
    if len(t) > 0 {
        tm = t[0]
    } else {
        tm = time.Now()
    }
    return int(tm.Month())
}

