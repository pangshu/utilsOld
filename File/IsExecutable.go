package File

import "os"

// IsExecutable 是否可执行文件.
func (*File)IsExecutable(filePath string) bool {
	info, err := os.Stat(filePath)
	return err == nil && info.Mode().IsRegular() && (info.Mode()&0111) != 0
}
