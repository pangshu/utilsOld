package File

import (
	"io"
	"os"
	"path/filepath"
)

// FastCopy 快速拷贝源文件到目标文件,不做安全检查.
func (toolFile *File)FastCopy(source string, dest string) (int64, error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = sourceFile.Close()
	}()

	//创建目录
	destDir := filepath.Dir(dest)
	if !toolFile.IsDir(destDir) {
		if err = os.MkdirAll(destDir, 0766); err != nil {
			return 0, err
		}
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = destFile.Close()
	}()

	var bufferSize int = 32768
	var nBytes int
	buf := make([]byte, bufferSize)
	for {
		n, err := sourceFile.Read(buf)
		if err != nil && err != io.EOF {
			return int64(nBytes), err
		} else if n == 0 {
			break
		}

		if _, err := destFile.Write(buf[:n]); err != nil || !toolFile.IsExist(dest) {
			return int64(nBytes), err
		}

		nBytes += n
	}

	return int64(nBytes), err
}
