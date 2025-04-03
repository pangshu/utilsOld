package File

import (
	"path/filepath"
	"regexp"
	"strings"
)

// FormatPath 格式化路径
func (toolFile *File)FormatPath(filePath string) string {
	if filePath == "" {
		return ""
	}

	//替换特殊字符
	filePath = strings.NewReplacer(`|`, "", `:`, "", `<`, "", `>`, "", `?`, "", `\`, "/").Replace(filePath)

	//替换连续斜杠, 连续的"//"或"\\"或"\/"或"/\"
	RegFormatDir := regexp.MustCompile(`[\/]{2,}`)

	filePath = RegFormatDir.ReplaceAllString(filePath, "/")

	dir := toolFile.FormatDir(filepath.Dir(filePath))

	if dir == `./` && strings.Index(filePath, dir) != -1 {
		return filePath
	}

	return dir + filepath.Base(filePath)
}