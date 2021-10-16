package String

import "unicode/utf8"

// IsUtf8 字符串是否UTF-8编码.
func (*String)IsUtf8(str string) bool {
	return str != "" && utf8.ValidString(str)
}

