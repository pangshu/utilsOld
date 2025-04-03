package Url

import (
	"net/url"
	"strings"
)

// Domain 从URL字符串中获取域名.
// 可选参数isMains,默认为false,取完整域名;为true时,取主域名(如abc.test.com取test.com).
func (toolUrl *Url)Domain(str string, isMain ...bool) string {
	str = toolUrl.FormatUrl(str)
	u, err := url.Parse(str)
	isMainType := false
	if len(isMain) > 0 {
		isMainType = isMain[0]
	}

	if err != nil {
		return ""
	} else if !isMainType {
		return u.Hostname()
	}

	parts := strings.Split(u.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]

	return domain
}

