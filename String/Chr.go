package String

// Chr 返回相对应于 ascii 所指定的单个字符.
func (*String)Chr(chr int) string {
	return string(rune(chr))
}
