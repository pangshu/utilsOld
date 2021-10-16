package String

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

// ToCamelCase 转为驼峰写法.
// 去掉包括下划线"_"和横杠"-".
func (*String)ToCamelCase(str string) string {
	if len(str) == 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	var r0, r1 rune
	var size int

	// leading connector will appear in output.
	for len(str) > 0 {
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		if !isCaseConnector(r0) {
			r0 = unicode.ToUpper(r0)
			break
		}

		buf.WriteRune(r0)
	}

	for len(str) > 0 {
		r1 = r0
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		if isCaseConnector(r0) && isCaseConnector(r1) {
			buf.WriteRune(r1)
			continue
		}

		if isCaseConnector(r1) {
			r0 = unicode.ToUpper(r0)
		} else if unicode.IsLower(r1) && unicode.IsUpper(r0) {
			buf.WriteRune(r1)
		} else if unicode.IsUpper(r1) && unicode.IsLower(r0) {
			buf.WriteRune(r1)
		} else {
			r0 = unicode.ToLower(r0)
			buf.WriteRune(r1)
		}
	}

	buf.WriteRune(r0)
	return buf.String()
}


// isCaseConnector 是否字符转换连接符.
func isCaseConnector(r rune) bool {
	return r == '-' || r == '_' || unicode.IsSpace(r)
}