package Number

import (
	"fmt"
	"strconv"
)

// Sum 对任意类型序列中的数值类型求和,忽略非数值的.
func (toolNumber *Number)Sum(nums []interface{}) (res float64) {
	if len(nums) == 0 {
		return
	} else {
		var err error
		var val float64
		for _, v := range nums {
			valueStr := fmt.Sprintf("%v", v)
			val, err = strconv.ParseFloat(valueStr, 64)
			if err == nil {
				res += val
			}
		}
	}

	return
}
