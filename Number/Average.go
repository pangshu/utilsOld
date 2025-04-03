package Number

import (
	"fmt"
	"strconv"
)

// Average 对数组序列求平均值.
func (*Number)Average(nums []interface{}) (res float64) {
	if len(nums) == 0 {
		return
	} else {
		var count int
		var err error
		var val, total float64
		for _, v := range nums {
			valueStr := fmt.Sprintf("%v", v)
			val, err = strconv.ParseFloat(valueStr, 64)
			if err == nil {
				count++
				total += val
			}
		}

		res = total / float64(count)
	}

	return
}
