package File

import "syscall"

// IsReadable 路径是否可读.
func (*File)IsReadable(filePath string) bool {
	err := syscall.Access(filePath, syscall.O_RDONLY)
	if err != nil {
		return false
	}
	return true
}
