package File

import "path/filepath"

// AbsPath 获取绝对路径,path可允许不存在.
func (*File)AbsPath(filePath string) string {
	fullPath := ""
	res, err := filepath.Abs(filePath) // filepath.Abs最终使用到os.Getwd()检查
	if err != nil {
		fullPath = filepath.Clean(filepath.Join(`/`, filePath))
	} else {
		fullPath = res
	}

	return fullPath
}
