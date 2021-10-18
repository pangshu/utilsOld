package Url

import "net/url"

// Decode 解码已编码的 URL 字符串.
func (*Url)Decode(str string) (string, error) {
	return url.QueryUnescape(str)
}
