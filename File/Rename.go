package File

import "os"

// Rename 重命名文件或目录.
func (*File)Rename(oldName, newName string) error {
	return os.Rename(oldName, newName)
}
