package Convert

import (
	"errors"
	"strings"
)

// JsonpToJson 将jsonp转为json串.  Example: forbar({a:"1",b:2}) to {"a":"1","b":2}
func (*Convert)JsonpToJson(str string) (string, error) {
	start := strings.Index(str, "(")
	end := strings.LastIndex(str, ")")

	if start == -1 || end == -1 {
		return "", errors.New("invalid jsonp.")
	}

	start += 1
	if start >= end {
		return "", errors.New("invalid jsonp.")
	}

	res := str[start:end]

	return res, nil
}
