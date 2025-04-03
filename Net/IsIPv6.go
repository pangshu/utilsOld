package Net

import (
	"net"
	"strings"
)

// IsIPv6 检查字符串是否IPv6地址.
func (*Net)IsIPv6(str string) bool {
	ipAddr := net.ParseIP(str)
	return ipAddr != nil && strings.ContainsRune(str, ':')
}
