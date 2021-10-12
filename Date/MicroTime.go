package Date

import "time"

// MicroTime 获取当前Unix时间戳(微秒).
func (*Date)MicroTime() int64 {
    return time.Now().UnixNano() / int64(time.Microsecond)
}