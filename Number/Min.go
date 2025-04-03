package Number

import (
	"fmt"
	"math"
	"strconv"
)

// Min 取出任意类型中数值类型的最小值,无数值类型则为0.
func (toolNumber *Number)Min(nums []interface{}) float64 {
	if len(nums) == 0 {
		return 0
	} else {
		var err error
		var val, res float64
		for _, v := range nums {
			valueStr := fmt.Sprintf("%v", v)
			val, err = strconv.ParseFloat(valueStr, 64)
			if err == nil {
				res = math.Min(res, val)
			}
		}
		return res
	}
}