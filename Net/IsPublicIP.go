package Net

import "net"

// IsPublicIP 是否公网IP.
func (*Net)IsPublicIP(ip string) bool {
	ips := net.ParseIP(ip)
	if ips.IsLoopback() || ips.IsLinkLocalMulticast() || ips.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := ips.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}
