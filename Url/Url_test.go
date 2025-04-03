package Url

import (
    "fmt"
    "net/url"
    "strings"
    "testing"
)

// 测试命令: go test -v -run TestBuildQuery Url/Global.go Url/BuildQuery.go Url/BuildQuery_test.go
func TestBuildQuery(t *testing.T) {
    var toolUrl Url
    params := url.Values{}
    params.Add("a", "abc")
    params.Add("b", "123")
    params.Add("c", "你好")
    res1 := toolUrl.BuildQuery(params)
    fmt.Println(res1)
}

// go test -v -run TestBuildQuery -bench=BenchmarkBuildQuery -count=5 Url/Global.go Url/BuildQuery.go Url/BuildQuery_test.go
func BenchmarkBuildQuery(t *testing.B) {
    t.ResetTimer()
    params := url.Values{}
    params.Add("a", "abc")
    params.Add("b", "123")
    params.Add("c", "你好")
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _ = toolUrl.BuildQuery(params)
    }
}

// 测试命令: go test -v -run TestClearPrefix Url/Global.go Url/ClearPrefix.go Url/ClearPrefix_test.go
func TestClearPrefix(t *testing.T) {
    var toolUrl Url
    str := "https://github.com/abc"
    arr := "https://"
    res1 := toolUrl.ClearPrefix(str,arr)
    fmt.Println(res1)
}

// go test -v -run TestClearPrefix -bench=BenchmarkClearPrefix -count=5 Url/Global.go Url/ClearPrefix.go Url/ClearPrefix_test.go
func BenchmarkClearPrefix(t *testing.B) {
    t.ResetTimer()
    str := "https://github.com/abc"
    arr := "https://"
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _ = toolUrl.ClearPrefix(str,arr)
    }
}

// 测试命令: go test -v -run TestClearSuffix Url/Global.go Url/ClearSuffix.go Url/ClearSuffix_test.go
func TestClearSuffix(t *testing.T) {
    var toolUrl Url
    str := "https://github.com/abc"
    arr := "com/abc"
    res1 := toolUrl.ClearSuffix(str,arr)
    fmt.Println(res1)
}

// go test -v -run TestClearSuffix -bench=BenchmarkClearSuffix -count=5 Url/Global.go Url/ClearSuffix.go Url/ClearSuffix_test.go
func BenchmarkClearSuffix(t *testing.B) {
    t.ResetTimer()
    str := "https://github.com/abc"
    arr := "com/abc"
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _ = toolUrl.ClearSuffix(str,arr)
    }
}

// 测试命令: go test -v -run TestDecode Url/Global.go Url/Decode.go Url/Decode_test.go
func TestDecode(t *testing.T) {
    var toolUrl Url
    str := "one%20%26%20two"
    res1,_ := toolUrl.Decode(str)
    fmt.Println(res1)
}

// go test -v -run TestDecode -bench=BenchmarkDecode -count=5 Url/Global.go Url/Decode.go Url/Decode_test.go
func BenchmarkDecode(t *testing.B) {
    t.ResetTimer()
    str := "one%20%26%20two"
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _,_ = toolUrl.Decode(str)
    }
}

// 测试命令: go test -v -run TestDomain Url/Global.go Url/Domain.go Url/Domain_test.go Url/FormatUrl.go
func TestDomain(t *testing.T) {
    var toolUrl Url
    str := "http://login.localhost:3000"
    res1 := toolUrl.Domain(str)
    fmt.Println(res1)
}

// go test -v -run TestDomain -bench=BenchmarkDomain -count=5 Url/Global.go Url/Domain.go Url/Domain_test.go
func BenchmarkDomain(t *testing.B) {
    t.ResetTimer()
    str := "http://login.localhost:3000"
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _ = toolUrl.Domain(str)
    }
}

// 测试命令: go test -v -run TestEncode Url/Global.go Url/Encode.go Url/Encode_test.go
func TestEncode(t *testing.T) {
    var toolUrl Url
    str := "one & two"
    res1 := toolUrl.Encode(str)
    fmt.Println(res1)
}

// go test -v -run TestEncode -bench=BenchmarkEncode -count=5 Url/Global.go Url/Encode.go Url/Encode_test.go
func BenchmarkEncode(t *testing.B) {
    t.ResetTimer()
    str := "one & two"
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _ = toolUrl.Encode(str)
    }
}

// 测试命令: go test -v -run TestFormatUrl Url/FormatUrl_test.go Url/Global.go Url/FormatUrl.go
func TestFormatUrl(t *testing.T) {
    var toolUrl Url
    res1 := toolUrl.FormatUrl("")
    res2 := toolUrl.FormatUrl("abc.com",2)
    res3 := toolUrl.FormatUrl("abc.com/hello?a=1",1)
    res4 := toolUrl.FormatUrl(`www.aaa.abc.com:8080\/a//b/c///d\\12/33\.jpg`)
    res5 := toolUrl.FormatUrl("/abc.com/hello?a=1",1)
    res6 := toolUrl.FormatUrl("https://abc.com/hello?a=1",1)
    res7 := toolUrl.FormatUrl(`/a//b/c///d\\12/33.jpg`)
    if res1 != "" || res2 == "" || res3 == "" || strings.ContainsRune(res4, '\\') || res5 == "" || res6 == "" || res7 == "" {
        t.Errorf("FormatUrl fail")
    }
    fmt.Println(res1)
    fmt.Println(res2)
    fmt.Println(res3)
    fmt.Println(res4)
    fmt.Println(res5)
    fmt.Println(res6)
    fmt.Println(res7)
}

// go test -v -run TestFormatUrl -bench=BenchmarkFormatUrl -count=5 Url/FormatUrl_test.go Url/Global.go Url/FormatUrl.go
func BenchmarkFormatUrl(t *testing.B) {
    t.StopTimer()
    t.StartTimer()
    for i:=0; i< t.N; i++ {
        var toolUrl Url
        res1 := toolUrl.FormatUrl("")
        res2 := toolUrl.FormatUrl("abc.com",2)
        res3 := toolUrl.FormatUrl("abc.com/hello?a=1",1)
        res4 := toolUrl.FormatUrl(`www.aaa.abc.com:8080\/a//b/c///d\\12/33.jpg`)
        res5 := toolUrl.FormatUrl("/abc.com/hello?a=1",1)
        res6 := toolUrl.FormatUrl("https://abc.com/hello?a=1",1)
        res7 := toolUrl.FormatUrl(`/a//b/c///d\\12/33.jpg`)
        if res1 != "" || res2 == "" || res3 == "" || strings.ContainsRune(res4, '\\') || res5 == "" || res6 == "" || res7 == "" {
            t.Errorf("FormatUrl fail")
        }
    }
}


// 测试命令: go test -v -run TestIsUrl Url/Global.go Url/IsUrl.go Url/IsUrl_test.go
func TestIsUrl(t *testing.T) {
    var toolUrl Url
    str := "http://login.localhost:3000"
    res1 := toolUrl.IsUrl(str)
    fmt.Println(res1)
}

// go test -v -run TestIsUrl -bench=BenchmarkIsUrl -count=5 Url/Global.go Url/IsUrl.go Url/IsUrl_test.go
func BenchmarkIsUrl(t *testing.B) {
    t.ResetTimer()
    str := "http://login.localhost:3000"
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _ = toolUrl.IsUrl(str)
    }
}

// 测试命令: go test -v -run TestIsUrlExists Url/Global.go Url/IsUrlExists.go Url/IsUrlExists_test.go  Url/IsUrl.go
func TestIsUrlExists(t *testing.T) {
    var toolUrl Url
    str := "http://baidu.com"
    res1 := toolUrl.IsUrlExists(str)
    fmt.Println(res1)
}

// go test -v -run TestIsUrlExists -bench=BenchmarkIsUrlExists -count=5 Url/Global.go Url/IsUrlExists.go Url/IsUrlExists_test.go  Url/IsUrl.go
func BenchmarkIsUrlExists(t *testing.B) {
    t.ResetTimer()
    str := "http://login.localhost:3000"
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _ = toolUrl.IsUrlExists(str)
    }
}

// 测试命令: go test -v -run TestParseStr Url/Global.go Url/ParseStr.go Url/ParseStr_test.go
func TestParseStr(t *testing.T) {
    var toolUrl Url
    str := `f[a][a]=m&f[a][b]=n`	//f[a][]=1&f[a][]=c&f[a][]=&f[b][]=bb&f[]=3&f[]=4
    arr1 := make(map[string]interface{})
    res1 := toolUrl.ParseStr(str, arr1)
    fmt.Println(res1)
}

// go test -v -run TestParseStr -bench=BenchmarkParseStr -count=5 Url/Global.go Url/ParseStr.go Url/ParseStr_test.go  Url/IsUrl.go
func BenchmarkParseStr(t *testing.B) {
    t.ResetTimer()
    str := `first=value&arr[]=foo+bar&arr[]=baz`
    arr1 := make(map[string]interface{})
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _ = toolUrl.ParseStr(str, arr1)
    }
}


// 测试命令: go test -v -run TestParseUrl Url/Global.go Url/ParseUrl.go Url/ParseUrl_test.go
func TestParseUrl(t *testing.T) {
    var toolUrl Url
    str := `https://www.cnbeta.com/articles/tech/1076401.htm`
    res1,_ := toolUrl.ParseUrl(str, -1)
    fmt.Println(res1)
}

// go test -v -run TestParseUrl -bench=BenchmarkParseUrl -count=5 Url/Global.go Url/ParseUrl.go Url/ParseUrl_test.go  Url/IsUrl.go
func BenchmarkParseUrl(t *testing.B) {
    t.ResetTimer()
    str := `https://www.cnbeta.com/articles/tech/1076401.htm`
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _,_ = toolUrl.ParseUrl(str, -1)
    }
}

// 测试命令: go test -v -run TestRawDecode Url/Global.go Url/RawDecode.go Url/RawDecode_test.go
func TestRawDecode(t *testing.T) {
    var toolUrl Url
    str := `foo%20bar%40baz`
    res1,_ := toolUrl.RawDecode(str)
    fmt.Println(res1)
}

// go test -v -run TestRawDecode -bench=BenchmarkRawDecode -count=5 Url/Global.go Url/RawDecode.go Url/RawDecode_test.go  Url/IsUrl.go
func BenchmarkRawDecode(t *testing.B) {
    t.ResetTimer()
    str := `foo%20bar%40baz`
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _,_ = toolUrl.RawDecode(str)
    }
}

// 测试命令: go test -v -run TestRawEncode Url/Global.go Url/RawEncode.go Url/RawEncode_test.go
func TestRawEncode(t *testing.T) {
    var toolUrl Url
    str := `foo bar@baz`
    res1 := toolUrl.RawEncode(str)
    fmt.Println(res1)
}

// go test -v -run TestRawEncode -bench=BenchmarkRawEncode -count=5 Url/Global.go Url/RawEncode.go Url/RawEncode_test.go  Url/IsUrl.go
func BenchmarkRawEncode(t *testing.B) {
    t.ResetTimer()
    str := `foo bar@baz`
    var toolUrl Url
    for i:=0; i< t.N; i++ {
        _ = toolUrl.RawEncode(str)
    }
}