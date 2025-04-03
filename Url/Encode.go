package Url

import "net/url"

// Encode 编码 URL 字符串.
func (*Url)Encode(str string) string {
	return url.QueryEscape(str)
}
