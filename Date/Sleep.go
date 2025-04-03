package Date

import "time"

// Sleep 延缓执行,秒.
func (*Date)Sleep(t int64) {
    time.Sleep(time.Duration(t) * time.Second)
}