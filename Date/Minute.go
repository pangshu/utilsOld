package Date

import "time"

// Minute 获取分钟.
func (*Date)Minute(t ...time.Time) int {
    var tm time.Time
    if len(t) > 0 {
        tm = t[0]
    } else {
        tm = time.Now()
    }
    return tm.Minute()
}
