package Number

import (
	"fmt"
	"strconv"
)

// IsBetween 数值是否在某个范围内,将自动转换类型再比较.
func (*Number)IsBetween(value interface{}, left interface{}, right interface{}) bool {
	valueStr := fmt.Sprintf("%v", value)
	leftStr := fmt.Sprintf("%v", left)
	rightStr := fmt.Sprintf("%v", right)

	valueNum,valueErr := strconv.ParseFloat(valueStr,64)
	if valueErr != nil {
		return false
	}
	leftNum,leftErr := strconv.ParseFloat(leftStr,64)
	if leftErr != nil {
		return false
	}
	rightNum,rightErr := strconv.ParseFloat(rightStr,64)
	if rightErr != nil {
		return false
	}

	if leftNum > rightNum {
		leftNum, rightNum = rightNum, leftNum
	}
	return valueNum >= leftNum && valueNum <= rightNum
}
