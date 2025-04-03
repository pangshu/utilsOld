package String

import "math"

const (
	PadLeft = "LEFT"
	PadRight = "RIGHT"
	PadBoth = "ALL"
)
// StrPad 使用fill填充str字符串到指定长度max.
// pType为填充类型,枚举值(PAD_LEFT,PAD_RIGHT,PAD_BOTH).
func (*String)StrPad(str string, fill string, max int, pType ...string) string {
	runeStr := []rune(str)
	runeStrLen := len(runeStr)
	if runeStrLen >= max || max < 1 || len(fill) == 0 {
		return str
	}

	padType := PadBoth
	if len(pType) > 0 {
		if pType[0] == PadLeft || pType[0] == PadRight || pType[0] == PadBoth {
			padType = pType[0]
		}
	}


	var leftSize int
	var rightSize int

	switch padType {
	case PadBoth:
		tmpSize := float64(max-runeStrLen) / 2
		leftSize = int(tmpSize)
		rightSize = int(tmpSize + math.Copysign(0.5, tmpSize))

	case PadLeft:
		leftSize = max - runeStrLen

	case PadRight:
		rightSize = max - runeStrLen
	}

	buf := make([]rune, 0, max)

	if padType == PadLeft || padType == PadBoth {
		for i := 0; i < leftSize; {
			for _, v := range []rune(fill) {
				buf = append(buf, v)
				if i >= leftSize-1 {
					i = leftSize
					break
				} else {
					i++
				}
			}
		}
	}

	buf = append(buf, runeStr...)

	if padType == PadRight || padType == PadBoth {
		for i := 0; i < rightSize; {
			for _, v := range []rune(fill) {
				buf = append(buf, v)
				if i >= rightSize-1 {
					i = rightSize
					break
				} else {
					i++
				}
			}
		}
	}

	return string(buf)
}

