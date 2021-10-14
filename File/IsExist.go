package File

import "os"

// IsExist 路径(文件/目录)是否存在.
func (*File)IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}
