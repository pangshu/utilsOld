package Date

import "time"

// Hour 获取小时.
func (*Date)Hour(t ...time.Time) int {
    var tm time.Time
    if len(t) > 0 {
        tm = t[0]
    } else {
        tm = time.Now()
    }
    return tm.Hour()
}

