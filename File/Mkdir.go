package File

import "os"

// Mkdir 新建目录,允许多级目录.
func (*File)Mkdir(filePath string, mode os.FileMode) error {
	return os.MkdirAll(filePath, mode)
}
