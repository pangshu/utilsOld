package Html

import "html"

// Htmlentities 将字符转换为 HTML 转义字符.
func (*Html)Htmlentities(str string) string {
	return html.EscapeString(str)
}
