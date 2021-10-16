package String

import "strings"

// TrimRight 删除字符串末端的空白字符（或者其他字符）.
func (*String)TrimRight(str string, characterMask ...string) string {
	//mask := getTrimMask(characterMask)
	var mask string
	if len(characterMask) == 0 {
		mask = " \t\n\r\v\f\x00　"
	} else {
		mask = strings.Join(characterMask, "")
	}
	return strings.TrimRight(str, mask)
}
