package String

import "strings"

// CountBase64Byte 统计base64字符串大小,字节.
func (*String)CountBase64Byte(str string) (res int) {
	pos := strings.Index(str, ",")
	if pos > 10 {
		img := strings.Replace(str[pos:], "=", "", -1)
		res = int(float64(len(img)) * float64(3.0/4.0))
	}

	return
}
