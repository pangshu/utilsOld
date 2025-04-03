package File

import (
	"os"
)

// IsReadable 路径是否可读.
func (*File) IsReadable(filePath string) bool {
	//err := syscall.Access(filePath, syscall.O_RDONLY)
	//if err != nil {
	//	return false
	//}

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return false
			//log.Println("Error: Read permission denied.")
		}
	}
	closeErr := file.Close()
	if closeErr != nil {
		return false
	} else {
		return true
	}
}
