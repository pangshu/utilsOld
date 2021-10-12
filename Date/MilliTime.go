package Date

import "time"

// MilliTime 获取当前Unix时间戳(毫秒).
func (*Date)MilliTime() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}