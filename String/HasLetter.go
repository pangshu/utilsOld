package String

// HasLetter 字符串是否含有(英文)字母.
func (*String)HasLetter(str string) bool {
	for _, r := range str {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return true
		}
	}
	return false
}
