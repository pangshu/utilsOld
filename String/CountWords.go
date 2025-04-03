package String

import (
	"bytes"
	"strings"
	"unicode"
)

// CountWords 统计字符串中单词的使用情况.
// 返回结果:单词总数;和一个字典,包含每个单词的单独统计.
// 因为没有分词,对中文尚未很好支持.
func (*String)CountWords(str string) (int, map[string]int) {
	//过滤标点符号
	var buffer bytes.Buffer
	for _, r := range str {
		if unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsMark(r) {
			buffer.WriteRune(' ')
		} else {
			buffer.WriteRune(r)
		}
	}

	var total int
	mp := make(map[string]int)
	words := strings.Fields(buffer.String())
	for _, word := range words {
		mp[word] += 1
		total++
	}

	return total, mp
}

