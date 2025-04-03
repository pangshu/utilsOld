package String

import "regexp"

// HasWhitespace 是否带有空白字符.
func (*String)HasWhitespace(str string) bool {
	//regexp.MustCompile(".*[[:space:]]").MatchString(str)
	return str != "" && regexp.MustCompile(".*[[:space:]]").MatchString(str)
}
