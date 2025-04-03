package String

// IsASCII 是否IsASCII字符串.
func (*String)IsASCII(str string) bool {
	//return str != "" && RegAscii.MatchString(str)
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] > 127 {
			return false
		}
	}

	return str != ""
}
