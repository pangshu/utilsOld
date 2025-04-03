package Net

import "net"

// LocalIp 提取本机所有IP
func (*Net)LocalIp() []string {
	res := make([]string, 0)
	addr, _ := net.InterfaceAddrs()
	if len(addr) > 0 {
		for _, addr := range addr {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if nil != ipNet.IP.To4() {
					res = append(res, ipNet.IP.String())
					//break
				}
			}
		}
	}
	return res
}