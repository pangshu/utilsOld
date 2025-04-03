package File

import (
	"os"
	"path/filepath"
)

// Chmod 批量改变路径权限模式(包括子目录和所属文件).fileMode为文件权限模式,dirMode为目录权限模式.
func (*File)Chmod(filePath string, fileMode os.FileMode) (res bool) {
	var err error
	err = filepath.Walk(filePath, func(filePath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		return os.Chmod(filePath, fileMode)
	})

	if err == nil {
		res = true
	}

	return
}
