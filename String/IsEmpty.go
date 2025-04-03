package String

// IsEmpty 字符串是否为空(包括空格).
func (toolString *String)IsEmpty(str string) bool {
	if str == "" || len(toolString.Trim(str)) == 0 {
		return true
	}

	return false
}
