package File

import (
	"bufio"
	"io"
	"os"
)

// ReadLastLine 读取文件末行.
func (*File)ReadLastLine(filePath string) string {
	var res string
	file, err := os.Open(filePath)
	if err == nil {
		var lastLineSize int
		reader := bufio.NewReader(file)

		for {
			bs, err := reader.ReadBytes('\n')
			lastLineSize = len(bs)
			if err == io.EOF {
				break
			}
		}

		fileInfo, _ := os.Stat(filePath)

		// make a buffer size according to the lastLineSize
		buffer := make([]byte, lastLineSize)
		offset := fileInfo.Size() - int64(lastLineSize)
		numRead, _ := file.ReadAt(buffer, offset)
		res = string(buffer[:numRead])
	}
	defer func() {
		_ = file.Close()
	}()

	return res
}

