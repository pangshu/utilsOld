package Url

import (
	"net/url"
	"strings"
)

// RawEncode 按照 RFC 3986 对 URL 进行编码.
func (*Url)RawEncode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}
