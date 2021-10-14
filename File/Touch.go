package File

import (
	"os"
	"path/filepath"
)

// Touch 快速创建指定大小的文件,size为字节.
func (toolFile *File)Touch(filePath string, size int64) bool {
	//创建目录
	destDir := filepath.Dir(filePath)
	if destDir != "" && !toolFile.IsDir(destDir) {
		if err := os.MkdirAll(destDir, 0766); err != nil {
			return false
		}
	}

	fd, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return false
	}
	defer func() {
		_ = fd.Close()
	}()

	if size > 1 {
		_, _ = fd.Seek(size-1, 0)
		_, _ = fd.Write([]byte{0})
	}

	return true
}
