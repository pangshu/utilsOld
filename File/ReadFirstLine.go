package File

import (
	"bufio"
	"os"
)

// ReadFirstLine 读取文件首行.
func (*File)ReadFirstLine(filePath string) string {
	var res string
	fileInfo, err := os.Open(filePath)
	if err == nil {
		scanner := bufio.NewScanner(fileInfo)
		for scanner.Scan() {
			res = scanner.Text()
			break
		}
	}
	defer func() {
		_ = fileInfo.Close()
	}()

	return res
}
