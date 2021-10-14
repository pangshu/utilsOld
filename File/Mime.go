package File

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

// Mime 获取文件mime类型;fast为true时根据后缀快速获取;为false时读取文件头获取.
func (*File)Mime(filePath string, fast bool) string {
	var res string
	if fast {
		suffix := filepath.Ext(filePath)
		res = mime.TypeByExtension(suffix)
	} else {
		srcFile, err := os.Open(filePath)
		if err != nil {
			return res
		}

		buffer := make([]byte, 512)
		_, err = srcFile.Read(buffer)
		if err != nil {
			return res
		}

		res = http.DetectContentType(buffer)
	}

	return res
}
