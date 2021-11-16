package Number

import (
	"errors"
	"math/rand"
	"time"
)

// RandArray 返回随机整型数组
func (*Number)RandArray(min, max, size int, isRepeat ...bool) ([]int, error) {
	var repeatStatus = true
	if len(isRepeat) > 0 {
		repeatStatus = isRepeat[0]
	}

	// 防止出错，对min与max进行互换
	if min > max {
		min, max = max, min
	}

	// 不重复模式下面，数组数量不能大于最大值
	if repeatStatus == false && (max < size || (max-min+1) < size) {
		return nil, errors.New("数据不能重复模式下面，数组数量不能大于最大值")
	}

	var slice []int
	for i := min; i <= max; i++ {
		slice = append(slice, i)
	}

	rand.Seed(time.Now().UnixNano())
	var ret []int
	if repeatStatus == false {
		// 数据不能重复模式
		for i := 0; i < size; i++ {
			idx := rand.Intn(len(slice))
			if slice != nil {
				ret = append(ret, slice[idx])
				slice = append(slice[:idx], slice[idx+1:]...)
			}

		}
	} else {
		// 数据可重复模式
		var lenSlice = len(slice)
		for i := 0; i < size; i++ {
			idx := rand.Intn(lenSlice)
			if slice != nil {
				ret = append(ret, slice[idx])
			}

		}
	}

	return ret, nil
}
