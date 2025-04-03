package Date

import "time"

// Year 获取年份.
func (*Date)Year(t ...time.Time) int {
    var tm time.Time
    if len(t) > 0 {
        tm = t[0]
    } else {
        tm = time.Now()
    }
    return tm.Year()
}
