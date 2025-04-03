package String

import "unicode/utf8"

// MbStrLen 获取宽字符串的长度,多字节的字符被计为 1.
func (*String)MbStrLen(str string) int {
	return utf8.RuneCountInString(str)
}
