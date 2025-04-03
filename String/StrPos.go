package String

import "strings"

// StrPos 查找字符串首次出现的位置,找不到时返回-1.
// haystack在该字符串中进行查找,needle要查找的字符串;
// offset起始位置,为负数时时,搜索会从字符串结尾指定字符数开始.
func (*String)StrPos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}
