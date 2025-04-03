package File

import (
	"encoding/base64"
	"io/ioutil"
	"strings"
)

// ToBase64 读取图片文件,并转换为base64字符串.
func (toolFile *File)ToBase64(filePath string) (string, string, error) {
	imgBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", "", err
	}

	fileType := strings.ToLower(toolFile.Ext(filePath))
	return base64.StdEncoding.EncodeToString(imgBuffer), fileType, nil
	//return fmt.Sprintf("data:image/%s;base64,%s", fileType, base64.StdEncoding.EncodeToString(imgBuffer)), fileType, nil
}
