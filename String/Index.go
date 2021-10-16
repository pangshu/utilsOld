package String

import "strings"

// Index 查找子串sub在字符串str中第一次出现的位置,不存在则返回-1;
// ignoreCase为是否忽略大小写.
func (*String)Index(str, sub string, ignoreCase ...bool) int {
	if str == "" || sub == "" {
		return -1
	}
	isLower := false
	if len(ignoreCase) > 0 {
		isLower = ignoreCase[0]
	}
	if isLower {
		str = strings.ToLower(str)
		sub = strings.ToLower(sub)
	}

	return strings.Index(str, sub)
}
