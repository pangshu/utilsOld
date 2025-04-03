package File

// IsImg 是否图片文件(仅检查后缀).
func (toolFile *File)IsImg(filePath string) bool {
	isBool := toolFile.IsExist(filePath)
	if true == isBool {
		ext := toolFile.Ext(filePath)
		switch ext {
		case "jpg", "jpeg", "bmp", "gif", "png", "svg", "ico", "webp":
			return true
		default:
			return false
		}
	} else {
		return false
	}
}
