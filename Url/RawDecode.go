package Url

import (
	"net/url"
	"strings"
)

// RawDecode 对已编码的 URL 字符串进行解码.
func (*Url)RawDecode(str string) (string, error) {
	return url.QueryUnescape(strings.Replace(str, "%20", "+", -1))
}
