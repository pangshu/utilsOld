package File

import (
	"path"
	"regexp"
	"strings"
)

// SafeFileName 将文件名转换为安全可用的字符串.
func (*File)SafeFileName(str string) string {
	name := strings.ToLower(str)
	name = path.Clean(path.Base(name))
	name = strings.Trim(name, " ")
	separators, err := regexp.Compile(`[ &_=+:]`)
	if err == nil {
		name = separators.ReplaceAllString(name, "-")
	}
	legal, err := regexp.Compile(`[^[:alnum:]-.]`)
	if err == nil {
		name = legal.ReplaceAllString(name, "")
	}
	for strings.Contains(name, "--") {
		name = strings.Replace(name, "--", "-", -1)
	}
	return name
}

