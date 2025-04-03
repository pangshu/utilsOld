package String

import "strings"

// Dstrpos 检查字符串str是否包含数组arr的元素之一,返回检查结果和匹配的字符串.
// chkCase为是否检查大小写.
func (toolString *String)Dstrpos(str string, arr []string, chkCase bool) (bool, string) {
	if len(str) == 0 || len(arr) == 0 {
		return false, ""
	}

	for _, v := range arr {
		var tmpStr string
		var valStr string
		if chkCase == false {
			tmpStr = strings.ToLower(str)
			valStr = strings.ToLower(v)
		}
		if toolString.StrPos(tmpStr, valStr, 0) != -1 {
			return true, valStr
		}
	}

	return false, ""
}
