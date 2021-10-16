package String

import (
	"fmt"
	"unicode/utf8"
)

// Ord 将首字符转换为rune.
func (*String)Ord(char string) rune {
	r, a := utf8.DecodeRune([]byte(char))
	fmt.Println(r)
	fmt.Println(a)
	return r
}
