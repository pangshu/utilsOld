package File

import "path/filepath"

// Dirname 返回路径中的目录部分,注意空路径或无目录的返回"." .
func (toolFile *File)Dirname(filePath string) string {
	if toolFile.IsDir(filePath) {
		return filePath
	} else {
		return filepath.Dir(filePath)
	}
}

