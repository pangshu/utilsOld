package Html

import "strings"

// NlToBr 将换行符转换为br标签.
func (*Html)NlToBr(html string) string {
	if html == "" {
		return ""
	}
	return strings.Replace(html, "\n", "<br />", -1)
}
