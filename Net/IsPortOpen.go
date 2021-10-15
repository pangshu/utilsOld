package Net

import (
	"net"
	"strconv"
	"strings"
	"time"
)

// IsPortOpen 检查主机端口是否开放.protocols为协议名称,可选,默认tcp.
func (*Net)IsPortOpen(host string, port int64, protocols ...string) bool {
	// 默认tcp协议
	protocol := "tcp"
	if len(protocols) > 0 && len(protocols[0]) > 0 {
		protocol = strings.ToLower(protocols[0])
	}

	conn, _ := net.DialTimeout(protocol, net.JoinHostPort(host, strconv.FormatInt(port, 10)), time.Millisecond*time.Duration(200))
	if conn != nil {
		_ = conn.Close()
		return true
	}

	return false
}











//package main
//
//import (
//"fmt"
//"net"
//"os"
//"sync"
//"time"
//)
//
//func checkPort(ip net.IP, port int,wg *sync.WaitGroup) {
//	tcpAddr := net.TCPAddr{
//		IP:   ip,
//		Port: port,
//	}
//	ch := make(chan bool)
//	timeout := make(chan bool)
//	go func(){
//		time.Sleep(3*time.Second)
//		timeout <- true
//	}()
//	go func(){
//		conn, err := net.DialTCP("tcp", nil, &tcpAddr)
//		ch <- true
//		if err == nil {
//			fmt.Printf("ip: %v port: %v \n",ip,port)
//			defer func(){
//				if conn != nil{
//					e := conn.Close()
//					if e !=nil{
//						fmt.Println(e)
//					}
//				}
//			}()
//		}
//	}()
//	select {
//	case <- timeout:
//		wg.Done()
//	case <- ch:
//		wg.Done()
//	}
//}
//func checkIp(ip string) bool{
//	if net.ParseIP(ip) == nil {
//		fmt.Println("非法ip地址")
//		return false
//	}else {
//		return true
//	}
//}
//func main() {
//
//	startTime := time.Now()
//	wg := sync.WaitGroup{}
//	wg.Add(65534)
//	ip := os.Args[1]
//	if checkIp(ip) {
//		for port := 1; port <= 65534;port++  {
//			go checkPort(net.ParseIP(ip),port,&wg)
//		}
//	}
//	wg.Wait()
//	endTime := time.Now()
//	fmt.Printf("执行时间%v",endTime.Sub(startTime))
//}