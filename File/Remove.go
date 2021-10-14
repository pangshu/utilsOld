package File

import "os"

// Remove 删除文件.
func (*File)Remove(filePath string) error {
	return os.Remove(filePath)
}