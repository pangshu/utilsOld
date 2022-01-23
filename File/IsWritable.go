package File

import (
	"log"
	"os"
)

// IsWritable 路径是否可写.
func (*File) IsWritable(filePath string) bool {
	//err := syscall.Access(filePath, syscall.O_RDWR)
	//if err != nil {
	//	return false
	//}

	file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	closeErr := file.Close()
	if closeErr != nil {
		return false
	} else {
		return true
	}
}
