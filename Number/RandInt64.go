package Number

import (
	"math/rand"
	"time"
)

// RandInt64 随机区间值
func (*Number)RandInt64(min, max int64) int64 {
	// 防止出错，对min与max进行互换
	if min > max {
		min, max = max, min
	}

	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min+1)
}
