package Url

import (
	"strings"
)

// ClearSuffix 清除URL的后缀;
// str为URL字符串,suffix为后缀,默认"/".
func (*Url)ClearSuffix(str string, suffix ...string) string {
	s := "/"
	if len(suffix) > 0 {
		s = suffix[0]
	}
	for s != "" && strings.HasSuffix(str, s) {
		str = str[0 : len(str)-len(s)]
	}

	return str
}
