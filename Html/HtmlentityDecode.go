package Html

import "html"

// HtmlentityDecode 将HTML实体转换为它们对应的字符.
func HtmlentityDecode(str string) string {
	return html.UnescapeString(str)
}

