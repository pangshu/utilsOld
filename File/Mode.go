package File

import "os"

// Mode 获取路径的权限模式.
func (*File)Mode(filePath string) (os.FileMode, error) {
	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Mode(), nil
}
