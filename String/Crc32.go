package String

import "hash/crc32"

// Crc32 计算一个字符串的 crc32 多项式.
func (*String)Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
