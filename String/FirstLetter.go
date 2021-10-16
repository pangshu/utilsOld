package String

import (
	"bytes"
	"encoding/hex"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)


// FirstLetter 获取字符串首字母.
func (toolString *String)FirstLetter(str string) string {
	if str != "" {
		// 获取字符串第一个字符
		_, size := utf8.DecodeRuneInString(str)
		firstChar := str[:size]
		//
		//var hzRegexp, _ = regexp.Compile("([\u4e00-\u9fa5]+)")
		//fmt.Println(hzRegexp.MatchString(firstChar))

		if regexp.MustCompile("^[\u4e00-\u9fa5]+$").MatchString(firstChar) {
			// Utf8 转 GBK2312
			reader := transform.NewReader(bytes.NewReader([]byte(firstChar)), simplifiedchinese.GBK.NewEncoder())
			firstCharGbk, _ := ioutil.ReadAll(reader)

			// 获取第一个字符的16进制
			firstCharHex := hex.EncodeToString([]byte(firstCharGbk))

			// 16进制转十进制
			firstCharDec, _ := strconv.ParseInt(firstCharHex, 16, 0)

			// 十进制落在GB 2312的某个拼音区间即为某个字母
			firstCharDecimalRelative := firstCharDec - 65536
			if firstCharDecimalRelative >= -20319 && firstCharDecimalRelative <= -20284 {
				return "A"
			}
			if firstCharDecimalRelative >= -20283 && firstCharDecimalRelative <= -19776 {
				return "B"
			}
			if firstCharDecimalRelative >= -19775 && firstCharDecimalRelative <= -19219 {
				return "C"
			}
			if firstCharDecimalRelative >= -19218 && firstCharDecimalRelative <= -18711 {
				return "D"
			}
			if firstCharDecimalRelative >= -18710 && firstCharDecimalRelative <= -18527 {
				return "E"
			}
			if firstCharDecimalRelative >= -18526 && firstCharDecimalRelative <= -18240 {
				return "F"
			}
			if firstCharDecimalRelative >= -18239 && firstCharDecimalRelative <= -17923 {
				return "G"
			}
			if firstCharDecimalRelative >= -17922 && firstCharDecimalRelative <= -17418 {
				return "H"
			}
			if firstCharDecimalRelative >= -17417 && firstCharDecimalRelative <= -16475 {
				return "J"
			}
			if firstCharDecimalRelative >= -16474 && firstCharDecimalRelative <= -16213 {
				return "K"
			}
			if firstCharDecimalRelative >= -16212 && firstCharDecimalRelative <= -15641 {
				return "L"
			}
			if firstCharDecimalRelative >= -15640 && firstCharDecimalRelative <= -15166 {
				return "M"
			}
			if firstCharDecimalRelative >= -15165 && firstCharDecimalRelative <= -14923 {
				return "N"
			}
			if firstCharDecimalRelative >= -14922 && firstCharDecimalRelative <= -14915 {
				return "O"
			}
			if firstCharDecimalRelative >= -14914 && firstCharDecimalRelative <= -14631 {
				return "P"
			}
			if firstCharDecimalRelative >= -14630 && firstCharDecimalRelative <= -14150 {
				return "Q"
			}
			if firstCharDecimalRelative >= -14149 && firstCharDecimalRelative <= -14091 {
				return "R"
			}
			if firstCharDecimalRelative >= -14090 && firstCharDecimalRelative <= -13319 {
				return "S"
			}
			if firstCharDecimalRelative >= -13318 && firstCharDecimalRelative <= -12839 {
				return "T"
			}
			if firstCharDecimalRelative >= -12838 && firstCharDecimalRelative <= -12557 {
				return "W"
			}
			if firstCharDecimalRelative >= -12556 && firstCharDecimalRelative <= -11848 {
				return "X"
			}
			if firstCharDecimalRelative >= -11847 && firstCharDecimalRelative <= -11056 {
				return "Y"
			}
			if firstCharDecimalRelative >= -11055 && firstCharDecimalRelative <= -10247 {
				return "Z"
			}
		} else {
			return strings.ToUpper(firstChar)
		}
	}

	return ""
}
