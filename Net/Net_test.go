package Net

import (
    "fmt"
    "net/http"
    "testing"
)
///////////////////////////////////// 测试 ClientIp ///////////////////////////////////
// 测试命令: go test -v -run TestClientIp Net/*
func TestClientIp(t *testing.T) {
    var toolNet Net
    newRequest := func(remoteAddr, xRealIP string, xForwardedFor ...string) *http.Request {
        h := http.Header{}
        h.Set("X-Real-IP", xRealIP)
        for _, address := range xForwardedFor {
            h.Set("X-Forwarded-For", address)
        }

        return &http.Request{
            RemoteAddr: remoteAddr,
            Header:     h,
        }
    }

    c := newRequest("114.114.114.114", "8.8.8.8")
    res1 := toolNet.ClientIp(c)
    fmt.Println(res1)
}

// go test -v -run TestClientIp -bench=BenchmarkClientIp -count=5 Net/*
func BenchmarkClientIp(t *testing.B) {
    t.ResetTimer()
    var toolNet Net
    newRequest := func(remoteAddr, xRealIP string, xForwardedFor ...string) *http.Request {
        h := http.Header{}
        h.Set("X-Real-IP", xRealIP)
        for _, address := range xForwardedFor {
            h.Set("X-Forwarded-For", address)
        }

        return &http.Request{
            RemoteAddr: remoteAddr,
            Header:     h,
        }
    }

    for i:=0; i< t.N; i++ {
        c := newRequest("114.114.114.114", "8.8.8.8")
        _ = toolNet.ClientIp(c)
    }
}

///////////////////////////////////// 测试 IpToLong ///////////////////////////////////
// 测试命令: go test -v -run TestIpToLong Net/*
func TestIpToLong(t *testing.T) {
    var toolNet Net
    res1 := toolNet.IpToLong("114.114.114.114")
    fmt.Println(res1)
}

// go test -v -run TestIpToLong -bench=BenchmarkIpToLong -count=5 Net/*
func BenchmarkIpToLong(t *testing.B) {
    t.ResetTimer()
    var toolNet Net
    for i:=0; i< t.N; i++ {
        _ = toolNet.IpToLong("114.114.114.114")
    }
}

///////////////////////////////////// 测试 IsIP ///////////////////////////////////
// 测试命令: go test -v -run TestIsIP Net/*
func TestIsIP(t *testing.T) {
    var toolNet Net
    res1 := toolNet.IsIP("114.114.114.114")
    fmt.Println(res1)
}

// go test -v -run TestIsIP -bench=BenchmarkIsIP -count=5 Net/*
func BenchmarkIsIP(t *testing.B) {
    t.ResetTimer()
    var toolNet Net
    for i:=0; i< t.N; i++ {
        _ = toolNet.IsIP("114.114.114.114")
    }
}

///////////////////////////////////// 测试 IsIPv4 ///////////////////////////////////
// 测试命令: go test -v -run TestIsIPv4 Net/*
func TestIsIPv4(t *testing.T) {
    var toolNet Net
    res1 := toolNet.IsIPv4("114.114.114.114")
    fmt.Println(res1)
}

// go test -v -run TestIsIPv4 -bench=BenchmarkIsIPv4 -count=5 Net/*
func BenchmarkIsIPv4(t *testing.B) {
    t.ResetTimer()
    var toolNet Net
    for i:=0; i< t.N; i++ {
        _ = toolNet.IsIPv4("114.114.114.114")
    }
}

///////////////////////////////////// 测试 IsIPv6 ///////////////////////////////////
// 测试命令: go test -v -run TestIsIPv6 Net/*
func TestIsIPv6(t *testing.T) {
    var toolNet Net
    res1 := toolNet.IsIPv6("2001:A304:6101:1::E0:F726:4E58")
    fmt.Println(res1)
}

// go test -v -run TestIsIPv6 -bench=BenchmarkIsIPv6 -count=5 Net/*
func BenchmarkIsIPv6(t *testing.B) {
    t.ResetTimer()
    var toolNet Net
    for i:=0; i< t.N; i++ {
        _ = toolNet.IsIPv6("2001:A304:6101:1::E0:F726:4E58")
    }
}

///////////////////////////////////// 测试 IsMacAddr ///////////////////////////////////
// 测试命令: go test -v -run TestIsMacAddr Net/*
func TestIsMacAddr(t *testing.T) {
    var toolNet Net
    res1 := toolNet.IsMacAddr("2001:A304:6101:1::E0:F726:4E58")
    fmt.Println(res1)
}

// go test -v -run TestIsMacAddr -bench=BenchmarkIsMacAddr -count=5 Net/*
func BenchmarkIsMacAddr(t *testing.B) {
    t.ResetTimer()
    var toolNet Net
    for i:=0; i< t.N; i++ {
        _ = toolNet.IsMacAddr("2001:A304:6101:1::E0:F726:4E58")
    }
}

///////////////////////////////////// 测试 IsPortOpen ///////////////////////////////////
// 测试命令: go test -v -run TestIsPortOpen Net/*
func TestIsPortOpen(t *testing.T) {
    var toolNet Net
    res1 := toolNet.IsPortOpen("127.0.0.1", 80)
    fmt.Println(res1)
}

// go test -v -run TestIsPortOpen -bench=BenchmarkIsPortOpen -count=5 Net/*
func BenchmarkIsPortOpen(t *testing.B) {
    t.ResetTimer()
    var toolNet Net
    for i:=0; i< t.N; i++ {
        _ = toolNet.IsPortOpen("127.0.0.1", 80)
    }
}

///////////////////////////////////// 测试 IsPublicIP ///////////////////////////////////
// 测试命令: go test -v -run TestIsPublicIP Net/*
func TestIsPublicIP(t *testing.T) {
    var toolNet Net
    res1 := toolNet.IsPublicIP("192.168.0.1")
    fmt.Println(res1)
}

// go test -v -run TestIsPublicIP -bench=BenchmarkIsPublicIP -count=5 Net/*
func BenchmarkIsPublicIP(t *testing.B) {
    t.ResetTimer()
    var toolNet Net
    for i:=0; i< t.N; i++ {
        _ = toolNet.IsPublicIP("192.168.0.1")
    }
}

///////////////////////////////////// 测试 LocalIp ///////////////////////////////////
// 测试命令: go test -v -run TestLocalIp Net/*
func TestLocalIp(t *testing.T) {
    var toolNet Net
    res1 := toolNet.LocalIp()
    fmt.Println(res1)
}

// go test -v -run TestLocalIp -bench=BenchmarkLocalIp -count=5 Net/*
func BenchmarkLocalIp(t *testing.B) {
    t.ResetTimer()
    var toolNet Net
    for i:=0; i< t.N; i++ {
        _ = toolNet.LocalIp()
    }
}

///////////////////////////////////// 测试 LocalMac ///////////////////////////////////
// 测试命令: go test -v -run TestLocalMac Net/*
func TestLocalMac(t *testing.T) {
    var toolNet Net
    res1 := toolNet.LocalMac()
    fmt.Println(res1)
}

// go test -v -run TestLocalMac -bench=BenchmarkLocalMac -count=5 Net/*
func BenchmarkLocalMac(t *testing.B) {
    t.ResetTimer()
    var toolNet Net
    for i:=0; i< t.N; i++ {
        _ = toolNet.LocalMac()
    }
}