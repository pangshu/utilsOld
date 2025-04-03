package Convert

import (
	"fmt"
	"strconv"
)

// ToUnicode 字符串转HTML编码
func (*Convert)ToUnicode(text string) string {
	var encoded string
	for _, letter := range text {
		a, _ := strconv.ParseInt(fmt.Sprintf("%d", letter), 10, 32)
		u := fmt.Sprintf("%d", letter)

		if a > 127 {
			encoded += "&#" + string(u) + ";"
		} else {
			encoded += string(letter)
		}
	}
	return encoded
}
