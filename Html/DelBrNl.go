package Html

import "strings"

// DelBrNl 将换行符转换为br标签.
func (*Html)DelBrNl(html string) string {
	if html == "" {
		return ""
	}

	html = strings.Replace(html, "\n", "", -1)
	html = strings.Replace(html, "\r", "", -1)
	html = strings.Replace(html, "<br />", "", -1)
	return strings.Replace(html, "\n", "<br />", -1)
}
