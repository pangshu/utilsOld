package String

import "unicode"

// FirstToUpper 将字符串的第一个字符转换为大写.
func (*String)FirstToUpper(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}


//func FirstToUpper(str string) string {
//	if len(str) < 1 {
//		return ""
//	}
//	strArray := []rune(str)
//	if strArray[0] >= 97 && strArray[0] <= 122  {
//		strArray[0] = strArray[0] - 32
//	}
//	return string(strArray)
//}
