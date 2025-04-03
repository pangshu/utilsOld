package String

import (
	"regexp"
	"strings"
)

// DelSpace 移除字符串中的空白字符.
// all为true时移除全部空白,为false时只替换连续的空白字符为一个空格.
func (*String)DelSpace(str string, all bool) string {
	if all && str != "" {
		return strings.Join(strings.Fields(str), "")
	} else if str != "" {
		//先将2个以上的连续空白符转为空格
		str = regexp.MustCompile(`[[:space:]]{2,}|[\s\p{Zs}]{2,}`).ReplaceAllString(str, " ")
		//再将[\t\n\f\r]等转为空格
		str = regexp.MustCompile(`\s`).ReplaceAllString(str, " ")
	}

	return strings.TrimSpace(str)
}
