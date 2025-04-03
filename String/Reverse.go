package String

// Reverse 反转字符串.
func (*String)Reverse(str string) string {
	n := len(str)
	runes := make([]rune, n)
	for _, r := range str {
		n--
		runes[n] = r
	}
	return string(runes[n:])
}

