package File

// IsBinary 是否二进制文件(且存在).
func (toolFile *File)IsBinary(filePath string) bool {
	cont, err := toolFile.ReadFile(filePath)
	if err != nil {
		return false
	}

	for _, b := range string(cont) {
		if 0 == b {
			return true
		}
	}
	return false
}
