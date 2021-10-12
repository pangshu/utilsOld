package Date

import "time"

// Second 获取秒数.
func (*Date)Second(t ...time.Time) int {
    var tm time.Time
    if len(t) > 0 {
        tm = t[0]
    } else {
        tm = time.Now()
    }
    return tm.Second()
}
