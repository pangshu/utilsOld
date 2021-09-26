package Number

import (
	"math/rand"
	"time"
)

// RandInt 随机区间整数
func (*Number)RandInt(min, max int) int {
	// 防止出错，对min与max进行互换
	if min > max {
		min, max = max, min
	}

	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
