package Net

import (
	"net"
	"strings"
)

// IsIPv4 检查字符串是否IPv4地址.
func (*Net)IsIPv4(str string) bool {
	ipAddr := net.ParseIP(str)
	// 不是合法的IP地址
	if ipAddr == nil {
		return false
	}

	return ipAddr.To4() != nil && strings.ContainsRune(str, '.')
}
