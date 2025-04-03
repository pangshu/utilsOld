package Net

import "net"

// IsIP 检查字符串是否IP地址.
func (*Net)IsIP(str string) bool {
	return str != "" && net.ParseIP(str) != nil
}
