package String

import (
	"encoding/base64"
)

// ToBase64 将字符串转换为base64图片.ext为图片扩展名,默认jpg.
func (*String)ToBase64(content string) string {
	return base64.StdEncoding.EncodeToString([]byte(content))
}
