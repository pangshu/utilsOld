package String

// IsLetters 字符串是否全(英文)字母组成.
func (*String)IsLetters(str string) bool {
	for _, r := range str {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return str != ""
}
