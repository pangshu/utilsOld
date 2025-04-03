package Url

import (
	"regexp"
	"strings"
)
const (
	SslAll = 2
	SslYes = 1
	SslNo = 0
)
// FormatUrl 格式化URL.
func (*Url)FormatUrl(str string, urlType ...int) string {
	sslType := 0
	if len(urlType) > 0 {
		sslType = urlType[0]
	}
	if str != "" {
		//str = strings.ReplaceAll(str, "\\", "/")
		//str = regexp.MustCompile(`([^:])[\/]{2,}`).ReplaceAllString(str, "$1/")
		if strings.Index(str,"//") == -1 {
			switch sslType {
			case SslAll:
				str = "//" + str
			case SslYes:
				str = "https://" + str
			case SslNo:
				str = "http://" + str
			}
		}
		str = strings.ReplaceAll(str, "\\", "/")
		str = regexp.MustCompile(`([^:])[\/]{2,}`).ReplaceAllString(str, "$1/")
	}
	return str
}
