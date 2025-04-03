package String

import "unicode"

// HasSpecialChar 字符串是否含有特殊字符.
func (*String)HasSpecialChar(str string) (res bool) {
	if str == "" {
		return
	}

	for _, r := range str {
		// IsPunct 判断 r 是否为一个标点字符 (类别 P)
		// IsSymbol 判断 r 是否为一个符号字符
		// IsMark 判断 r 是否为一个 mark 字符 (类别 M)
		if unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsMark(r) {
			res = true
			return
		}
	}

	return
}

