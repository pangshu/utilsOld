package String

import "strings"

// SubstrCount 计算字符串出现的次数.
func (*String)SubstrCount(str, substr string) int {
	return strings.Count(str, substr)
}
