package String

import "strings"

// TrimLeft 删除字符串开头的空白字符（或其他字符）.
func (*String)TrimLeft(str string, characterMask ...string) string {
	var mask string
	if len(characterMask) == 0 {
		mask = " \t\n\r\v\f\x00　"
	} else {
		mask = strings.Join(characterMask, "")
	}
	return strings.TrimLeft(str, mask)
}
