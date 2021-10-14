package File

import "path/filepath"

// Basename 返回路径中的文件名部分.
func (*File)Basename(filePath string) string {
	return filepath.Base(filePath)
}
