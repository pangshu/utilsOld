package String

import "strings"

// Trim 去除字符串首尾处的空白字符（或者其他字符）.
func (*String)Trim(str string, characterMask ...string) string {
	//mask := getTrimMask(characterMask)
	var mask string
	if len(characterMask) == 0 {
		mask = " \t\n\r\v\f\x00　"
	} else {
		mask = strings.Join(characterMask, "")
	}
	return strings.Trim(str, mask)
}
