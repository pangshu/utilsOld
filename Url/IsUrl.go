package Url

import (
	"net/url"
	"strings"
	"unicode/utf8"
)

// IsUrl 检查字符串是否URL.
func (*Url)IsUrl(str string) bool {
	if str == "" || len(str) <= 3 || utf8.RuneCountInString(str) >= 2083 || strings.HasPrefix(str, ".") {
		return false
	}

	res, err := url.ParseRequestURI(str)
	if err != nil {
		return false //Couldn't even parse the url
	}
	if len(res.Scheme) == 0 {
		return false //No Scheme found
	}

	return true
}
