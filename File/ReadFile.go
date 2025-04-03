package File

import "io/ioutil"

// ReadFile 读取文件内容.
func (*File)ReadFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	return data, err
}
