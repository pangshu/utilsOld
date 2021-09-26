package Number

import (
	"fmt"
	"strconv"
)

// Format 以千位分隔符方式格式化一个数字. decimal为要保留的小数位数,point为小数点显示的字符,thousand为千位分隔符显示的字符.
func (*Number)Format(number float64, decimal int, point, thousand string) string {
	isMinus := false
	if number < 0 {
		number = -number
		isMinus = true
	}
	// 保留小数点
	decNum := decimal
	if decNum < 0 && isMinus == true{
		decNum = -decNum
	}

	numStr := fmt.Sprintf("%."+strconv.Itoa(decNum)+"F", number)
	if point == "" && thousand == "" {
		return numStr
	} else {
		prefix, suffix := "", ""
		if decNum > 0 {
			prefix = numStr[:len(numStr)-(decNum+1)]
			suffix = numStr[len(numStr)-decNum:]
		} else {
			prefix = numStr
		}
		sep := []byte(thousand)
		n, l1, l2 := 0, len(prefix), len(sep)
		// thousands sep num
		c := (l1 - 1) / 3
		tmp := make([]byte, l2*c+l1)
		pos := len(tmp) - 1
		for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
			if l2 > 0 && n > 0 && n%3 == 0 {
				for j := range sep {
					tmp[pos] = sep[l2-j-1]
					pos--
				}
			}
			tmp[pos] = prefix[i]
		}
		s := string(tmp)
		if decNum > 0 {
			s += point + suffix
		}
		if isMinus {
			s = "-" + s
		}

		return s
	}
}
