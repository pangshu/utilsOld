package File

import (
	"os"
	"path/filepath"
)

// RealPath 返回规范化的真实绝对路径名,path必须存在.若路径不存在则返回空字符串.
func (*File)RealPath(filePath string) string {
	fullPath := filePath
	if !filepath.IsAbs(filePath) {
		wd, err := os.Getwd()
		if err != nil {
			return ""
		}
		fullPath = filepath.Clean(wd + `/` + filePath)
	}

	_, err := os.Stat(fullPath)
	if err != nil {
		return ""
	}

	return fullPath
}
