package String

import "strings"

// Explode 字符串分割.delimiters为分隔符,可选,支持多个.
func (*String)Explode(str string, delimiters ...string) (res []string) {
	if str == "" {
		return
	}

	dLen := len(delimiters)
	if dLen == 0 {
		res = append(res, str)
	} else if dLen > 1 {
		var sl []string
		for _, v := range delimiters {
			if v != "" {
				sl = append(sl, v, "$@#KSYS#@$")
			}
		}
		str = strings.NewReplacer(sl...).Replace(str)
		res = strings.Split(str, "$@#KSYS#@$")
	} else {
		res = strings.Split(str, delimiters[0])
	}

	return
}
