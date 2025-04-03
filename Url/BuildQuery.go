package Url

import (
	"net/url"
)

// BuildQuery 根据参数生成 URL-encode 之后的请求字符串.
func (*Url)BuildQuery(queryData url.Values) string {
	return queryData.Encode()
}
