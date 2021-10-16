package String

import "unicode"

// HasChinese 字符串是否含有中文.
func (*String)HasChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}

	return false
}
