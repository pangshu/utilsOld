package File

import (
	"regexp"
	"strings"
)

// FormatDir 格式化目录,将"\","//"替换为"/",且以"/"结尾.
func (*File)FormatDir(filePath string) string {
	if filePath == "" {
		return ""
	}

	//替换特殊字符
	filePath = strings.NewReplacer(`|`, "", `:`, "", `<`, "", `>`, "", `?`, "", `\`, "/").Replace(filePath)
	//替换连续斜杠, 连续的"//"或"\\"或"\/"或"/\"
	RegFormatDir := regexp.MustCompile(`[\/]{2,}`)

	filePath = RegFormatDir.ReplaceAllString(filePath, "/")

	return strings.TrimRight(filePath, "/") + "/"
}
