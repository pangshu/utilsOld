package Number

import (
	"fmt"
	"strconv"
)

// Percent 返回百分比(val/total).
func (toolNumber *Number)Percent(val, total interface{}) float64 {
	valueStr := fmt.Sprintf("%v", val)
	totalStr := fmt.Sprintf("%v", total)

	valueFloat, valueErr := strconv.ParseFloat(valueStr, 64)
	if valueErr != nil {
		return 0
	}

	if valueFloat == 0 {
		return 0
	}

	totalFloat, totalErr := strconv.ParseFloat(totalStr, 64)
	if totalErr != nil {
		return 0
	}

	return (valueFloat / totalFloat) * 100
}
