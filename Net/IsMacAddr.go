package Net

import "net"

// IsMacAddr 是否MAC物理网卡地址.
func (*Net)IsMacAddr(str string) bool {
	_, err := net.ParseMAC(str)
	return err == nil
}
