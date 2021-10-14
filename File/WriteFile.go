package File

import (
	"io/ioutil"
	"os"
	"path"
)

// WriteFile 将内容写入文件.
// filePath为文件路径,data为内容,perm为权限.
func (*File)WriteFile(filePath string, data string, perm ...os.FileMode) error {
	if dir := path.Dir(filePath); dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	var p os.FileMode = 0655
	if len(perm) > 0 {
		p = perm[0]
	}

	return ioutil.WriteFile(filePath, []byte(data), p)
}
