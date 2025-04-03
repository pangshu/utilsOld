package Url

import "strings"

// ClearPrefix 清除URL的前缀;
// str为URL字符串,prefix为前缀,默认"/".
func (*Url)ClearPrefix(str string, prefix ...string) string {
	p := "/"
	if len(prefix) > 0 {
		p = prefix[0]
	}

	for p != "" && strings.HasPrefix(str, p) {
		str = str[len(p):len(str)]
	}

	return str
}
