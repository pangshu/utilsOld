package Net

import "net"

// LocalMac 获取本机MAC地址
func (*Net)LocalMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, inter := range interfaces {
		address, err := inter.Addrs()
		if err != nil {
			return ""
		}

		for _, address := range address {
			// check the address type and if it is not a loopback the display it
			if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					return inter.HardwareAddr.String()
				}
			}
		}
	}
	return ""
}
