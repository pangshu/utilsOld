package String

import "regexp"

// HasFullWidth 是否含有全角字符.
func (*String)HasFullWidth(str string) bool {
	//regexp.MustCompile("[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]").MatchString(str)
	return str != "" && regexp.MustCompile("[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]").MatchString(str)
}
