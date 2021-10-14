package File

import "os"

// IsDir 是否目录(且存在).
func (*File)IsDir(filePath string) bool {
	f, err := os.Lstat(filePath)
	if os.IsNotExist(err) || nil != err {
		return false
	}
	return f.IsDir()
}
