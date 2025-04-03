package Number

import (
    "math"
)

// Round 四舍五入
func (*Number)Round(num interface{}, decimal ...int) float64 {
    // 保留小数点
    decNum := 0
    if len(decimal) > 0 {
        decNum = decimal[0]
    }

    decPow := math.Pow10(decNum)
    numFloat := num.(float64)
    return float64(int(numFloat * decPow + 0.5)) / decPow
    //下面这种方案,运行比较慢
    //return math.Trunc((numFloat +0.5 / decPow) * decPow) / decPow
}
