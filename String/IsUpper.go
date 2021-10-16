package String

import "unicode"

// IsUpper 字符串是否全部大写.
func (*String)IsUpper(str string) bool {
	for _, r := range str {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return str != ""
}
