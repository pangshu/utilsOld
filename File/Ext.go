package File

import (
	"path/filepath"
	"strings"
)

// Ext 获取文件的小写扩展名,不包括点"." .
func (toolFile *File)Ext(filePath string) string {
	if toolFile.IsExist(filePath) {
		suffix := filepath.Ext(filePath)
		if suffix != "" {
			return strings.ToLower(suffix[1:])
		}
		return suffix
	} else {
		return ""
	}
}
