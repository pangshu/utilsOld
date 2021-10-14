package File

import "syscall"

// IsWritable 路径是否可写.
func (*File)IsWritable(filePath string) bool {
	err := syscall.Access(filePath, syscall.O_RDWR)
	if err != nil {
		return false
	}
	return true
}