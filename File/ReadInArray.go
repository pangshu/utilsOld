package File

import (
	"io/ioutil"
	"strings"
)

// ReadInArray 把整个文件读入一个数组中,每行作为一个元素.
func (*File)ReadInArray(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
