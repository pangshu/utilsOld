package File

import "os"

// IsLink 是否链接文件(且存在).
func (*File)IsLink(filePath string) bool {
	f, err := os.Lstat(filePath)
	if err != nil {
		return false
	}

	return f.Mode()&os.ModeSymlink == os.ModeSymlink
}
