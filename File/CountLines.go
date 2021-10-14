package File

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// CountLines 统计文件行数.buffLength为缓冲长度,kb.
func (*File)CountLines(filePath string, cacheLength ...int) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return -1, err
	}
	defer func() {
		_ = file.Close()
	}()

	count := 0
	lineSep := []byte{'\n'}

	buffLength := 32
	if len(cacheLength) > 0 {
		buffLength = cacheLength[0]
	}

	r := bufio.NewReader(file)
	buf := make([]byte, buffLength*1024)
	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}
