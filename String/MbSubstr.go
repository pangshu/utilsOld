package String

// MbSubstr 返回(宽字符)字符串str的子串.
// start 为起始位置.若值是负数,返回的结果将从 str 结尾处向前数第 abs(start) 个字符开始.
// length 为截取的长度.若值时负数, str 末尾处的 abs(length) 个字符将会被省略.
// start/length的绝对值必须<=原字符串长度.
func (*String)MbSubstr(str string, start int, length ...int) string {
	if str == "" {
		return ""
	}

	tmpStr := []rune(str)
	tmpStrLen := len(tmpStr)

	//起始位置，若start为负，则进行减计
	if start < 0 {
		start  = tmpStrLen + start
	}
	//结尾长度
	end := 0
	if len(length) == 0 {
		end = tmpStrLen - start
	} else {
		if length[0] + start > tmpStrLen {
			end = tmpStrLen - start
		} else {
			end = length[0]
		}
	}
	end = end + start
	return string(tmpStr[start:end])
}