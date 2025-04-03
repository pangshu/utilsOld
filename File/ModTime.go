package File

import (
	"os"
	"time"
)

// ModTime 获取文件的修改时间戳,秒.
func (*File)ModTime(filePath string) time.Time {
	fileInfo, err := os.Stat(filePath)
	if err == nil {
		return fileInfo.ModTime()
	} else {
		return time.Time{}
	}
}
