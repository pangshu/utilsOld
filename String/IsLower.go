package String

import "unicode"

// IsLower 字符串是否全部小写.
func (*String)IsLower(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return str != ""
}
