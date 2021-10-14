package File

import (
	"os"
	"path/filepath"
)

// DirSize 获取目录大小(bytes字节).
func (*File)DirSize(filePath string) int64 {
	var size int64
	//filepath.Walk压测很慢
	_ = filepath.Walk(filePath, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size
}
