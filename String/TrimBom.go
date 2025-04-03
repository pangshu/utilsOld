package String

import "bytes"

// TrimBom 移除字符串中的BOM
func (*String)TrimBom(str []byte) []byte {
	return bytes.Trim(str, "\xef\xbb\xbf")
}
