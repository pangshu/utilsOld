package Number

import "math"

// IsWhole 数值是否为整数.
func (*Number)IsWhole(value float64) bool {
	return math.Remainder(value, 1) == 0
}
