package String

import "regexp"

// StripTags 过滤html和php标签.
func (*String)StripTags(str string) string {
	if str == "" {
		return ""
	}

	return regexp.MustCompile(`<(.|\n)*?>`).ReplaceAllString(str, "")
}
