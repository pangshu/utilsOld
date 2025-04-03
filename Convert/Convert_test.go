package Convert

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"html/template"
	"io/ioutil"
	"reflect"
	"testing"
)


///////////////////////////////////// 测试 ToCharset ///////////////////////////////////
// 测试命令: go test -v -run TestToCharset Convert/*
func TestToCharset(t *testing.T) {
	var toolConvert Convert

	reader := transform.NewReader(bytes.NewReader([]byte("花间一壶酒，独酌无相亲。")), simplifiedchinese.GBK.NewEncoder())
	d, _ := ioutil.ReadAll(reader)
	fmt.Println(string(d))


	//var testData = []struct {
	//	utf8, other, otherEncoding string
	//}{
	//	{"Résumé", "Résumé", "utf-8"},
	//	//{"Résumé", "R\xe9sum\xe9", "latin-1"},
	//	{"これは漢字です。", "S0\x8c0o0\"oW[g0Y0\x020", "UTF-16LE"},
	//	{"これは漢字です。", "0S0\x8c0oo\"[W0g0Y0\x02", "UTF-16BE"},
	//	{"これは漢字です。", "\xfe\xff0S0\x8c0oo\"[W0g0Y0\x02", "UTF-16"},
	//	{"𝄢𝄞𝄪𝄫", "\xfe\xff\xd8\x34\xdd\x22\xd8\x34\xdd\x1e\xd8\x34\xdd\x2a\xd8\x34\xdd\x2b", "UTF-16"},
	//	//{"Hello, world", "Hello, world", "ASCII"},
	//	{"Gdańsk", "Gda\xf1sk", "ISO-8859-2"},
	//	{"Ââ Čč Đđ Ŋŋ Õõ Šš Žž Åå Ää", "\xc2\xe2 \xc8\xe8 \xa9\xb9 \xaf\xbf \xd5\xf5 \xaa\xba \xac\xbc \xc5\xe5 \xc4\xe4", "ISO-8859-10"},
	//	//{"สำหรับ", "\xca\xd3\xcb\xc3\u047a", "ISO-8859-11"},
	//	{"latviešu", "latvie\xf0u", "ISO-8859-13"},
	//	{"Seònaid", "Se\xf2naid", "ISO-8859-14"},
	//	{"€1 is cheap", "\xa41 is cheap", "ISO-8859-15"},
	//	{"românește", "rom\xe2ne\xbate", "ISO-8859-16"},
	//	{"nutraĵo", "nutra\xbco", "ISO-8859-3"},
	//	{"Kalâdlit", "Kal\xe2dlit", "ISO-8859-4"},
	//	{"русский", "\xe0\xe3\xe1\xe1\xda\xd8\xd9", "ISO-8859-5"},
	//	{"ελληνικά", "\xe5\xeb\xeb\xe7\xed\xe9\xea\xdc", "ISO-8859-7"},
	//	{"Kağan", "Ka\xf0an", "ISO-8859-9"},
	//	{"Résumé", "R\x8esum\x8e", "macintosh"},
	//	{"Gdańsk", "Gda\xf1sk", "windows-1250"},
	//	{"русский", "\xf0\xf3\xf1\xf1\xea\xe8\xe9", "windows-1251"},
	//	{"Résumé", "R\xe9sum\xe9", "windows-1252"},
	//	{"ελληνικά", "\xe5\xeb\xeb\xe7\xed\xe9\xea\xdc", "windows-1253"},
	//	{"Kağan", "Ka\xf0an", "windows-1254"},
	//	{"עִבְרִית", "\xf2\xc4\xe1\xc0\xf8\xc4\xe9\xfa", "windows-1255"},
	//	{"العربية", "\xc7\xe1\xda\xd1\xc8\xed\xc9", "windows-1256"},
	//	{"latviešu", "latvie\xf0u", "windows-1257"},
	//	{"Việt", "Vi\xea\xf2t", "windows-1258"},
	//	{"สำหรับ", "\xca\xd3\xcb\xc3\u047a", "windows-874"},
	//	{"русский", "\xd2\xd5\xd3\xd3\xcb\xc9\xca", "KOI8-R"},
	//	{"українська", "\xd5\xcb\xd2\xc1\xa7\xce\xd3\xd8\xcb\xc1", "KOI8-U"},
	//	{"Hello 常用國字標準字體表", "Hello \xb1`\xa5\u03b0\xea\xa6r\xbc\u0437\u01e6r\xc5\xe9\xaa\xed", "big5"},
	//	{"Hello 常用國字標準字體表", "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed", "gbk"},
	//	{"Hello 常用國字標準字體表", "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed", "gb18030"},
	//	{"花间一壶酒，独酌无相亲。", "~{;(<dR;:x>F#,6@WCN^O`GW!#", "GB2312"},
	//	{"花间一壶酒，独酌无相亲。", "~{;(<dR;:x>F#,6@WCN^O`GW!#", "HZGB2312"},
	//	{"עִבְרִית", "\x81\x30\xfb\x30\x81\x30\xf6\x34\x81\x30\xf9\x33\x81\x30\xf6\x30\x81\x30\xfb\x36\x81\x30\xf6\x34\x81\x30\xfa\x31\x81\x30\xfb\x38", "gb18030"},
	//	{"㧯", "\x82\x31\x89\x38", "gb18030"},
	//	{"㧯", "㧯", "UTF-8"},
	//	//{"これは漢字です。", "\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B", "SJIS"},
	//	{"これは漢字です。", "\xa4\xb3\xa4\xec\xa4\u03f4\xc1\xbb\xfa\xa4\u01e4\xb9\xa1\xa3", "EUC-JP"},
	//}


	var testData = []struct {
		utf8, other, otherEncoding string
	}{
		{"花间一壶酒，独酌无相亲。", "花间一壶酒，独酌无相亲。", "utf-8"},
	}
	for _, data := range testData {
		str := ""
		str, err := toolConvert.Charset(data.other, data.otherEncoding, "gbk")
		if err != nil {
			t.Errorf("Could not create decoder for %v", err)
			continue
		}

		if str != data.utf8 {
			t.Errorf("Unexpected value: %#v (expected %#v) %v", str, data.utf8, data.otherEncoding)
		}
		fmt.Println(str)
	}
}

// go test -v -run TestToStr -bench=BenchmarkToStr -count=5 Convert/*
//func BenchmarkToStr(t *testing.B) {
//	t.ResetTimer()
//	var toolConvert Convert
//	type Key struct {
//		k string
//	}
//	key := &Key{"foo"}
//
//	type Investment struct {
//		Price  float64
//		Symbol string
//		Rating int64
//	}
//	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}
//
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(8)},
//		{int8(8)},
//		{int16(8)},
//		{int32(8)},
//		{int64(8)},
//		{uint(8)},
//		{uint8(8)},
//		{uint16(8)},
//		{uint32(8)},
//		{uint64(8)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{nil},
//		{[]byte("one time")},
//		{"one more time"},
//		{template.HTML("one time")},
//		{template.URL("http://somehost.foo")},
//		{template.JS("(1+2)")},
//		{template.CSS("a")},
//		{template.HTMLAttr("a")},
//		{func() error {return nil}},
//		// errors
//		{testing.T{}},
//		{key},
//		{&inv.Price},
//		{&inv.Symbol},
//		{&inv.Rating},
//		{inv},
//	}
//	for i:=0; i< t.N; i++ {
//		for _, test := range tests {
//			_,_ = toolConvert.ToStr(test.input)
//		}
//
//	}
//}



///////////////////////////////////// 测试 ToStr ///////////////////////////////////
// 测试命令: go test -v -run TestToStr Convert/*
func TestToStr(t *testing.T) {
	var toolConvert Convert
	type Key struct {
		k string
	}
	key := &Key{"foo"}

	type Investment struct {
		Price  float64
		Symbol string
		Rating int64
	}
	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}

	tests := []struct {
		input  interface{}
	}{
		{int(8)},
		{int8(8)},
		{int16(8)},
		{int32(8)},
		{int64(8)},
		{uint(8)},
		{uint8(8)},
		{uint16(8)},
		{uint32(8)},
		{uint64(8)},
		{float32(8.31)},
		{float64(8.31)},
		{true},
		{false},
		{nil},
		{[]byte("one time")},
		{"one more time"},
		{template.HTML("one time")},
		{template.URL("http://somehost.foo")},
		{template.JS("(1+2)")},
		{template.CSS("a")},
		{template.HTMLAttr("a")},
		{func() error {return nil}},
		// errors
		{testing.T{}},
		{key},
		{&inv.Price},
		{&inv.Symbol},
		{&inv.Rating},
		{inv},
	}

	for _, test := range tests {
		v,err := toolConvert.ToStr(test.input)
		b := reflect.ValueOf(test.input)
		fmt.Println(test.input , " >>>>>>>>>>>> " , b.Kind() , " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
	}
}

// go test -v -run TestToStr -bench=BenchmarkToStr -count=5 Convert/*
func BenchmarkToStr(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	type Key struct {
		k string
	}
	key := &Key{"foo"}

	type Investment struct {
		Price  float64
		Symbol string
		Rating int64
	}
	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}

	tests := []struct {
		input  interface{}
	}{
		{int(8)},
		{int8(8)},
		{int16(8)},
		{int32(8)},
		{int64(8)},
		{uint(8)},
		{uint8(8)},
		{uint16(8)},
		{uint32(8)},
		{uint64(8)},
		{float32(8.31)},
		{float64(8.31)},
		{true},
		{false},
		{nil},
		{[]byte("one time")},
		{"one more time"},
		{template.HTML("one time")},
		{template.URL("http://somehost.foo")},
		{template.JS("(1+2)")},
		{template.CSS("a")},
		{template.HTMLAttr("a")},
		{func() error {return nil}},
		// errors
		{testing.T{}},
		{key},
		{&inv.Price},
		{&inv.Symbol},
		{&inv.Rating},
		{inv},
	}
	for i:=0; i< t.N; i++ {
		for _, test := range tests {
			_,_ = toolConvert.ToStr(test.input)
		}

	}
}

///////////////////////////////////// 测试 JsonpToJson ///////////////////////////////////
// 测试命令: go test -v -run TestJsonpToJson Convert/*
func TestJsonpToJson(t *testing.T) {
	var toolConvert Convert
	str := `myFunc([{"Name":"Bob","Age":32,"Company":"IBM","Engineer":true},{"Name":"John","Age":20,"Company":"Oracle","Engineer":false},{"Name":"Henry","Age":45,"Company":"Microsoft","Engineer":false}]);`
	res1,_ := toolConvert.JsonpToJson(str)
	fmt.Println(res1)
}

// go test -v -run TestJsonpToJson -bench=BenchmarkJsonpToJson -count=5 Convert/*
func BenchmarkJsonpToJson(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := `myFunc([{"Name":"Bob","Age":32,"Company":"IBM","Engineer":true},{"Name":"John","Age":20,"Company":"Oracle","Engineer":false},{"Name":"Henry","Age":45,"Company":"Microsoft","Engineer":false}]);`
	for i:=0; i< t.N; i++ {
		_,_ = toolConvert.JsonpToJson(str)
	}
}

///////////////////////////////////// 测试 JsonToMap ///////////////////////////////////
// 测试命令: go test -v -run TestJsonToMap Convert/*
func TestJsonToMap(t *testing.T) {
	var toolConvert Convert
	str := `{"Name":"Bob","Age":32,"Company":"IBM","Engineer":true}`
	res1,_ := toolConvert.JsonToMap(str)
	fmt.Println(res1)
}

// go test -v -run TestJsonToMap -bench=BenchmarkJsonToMap -count=5 Convert/*
func BenchmarkJsonToMap(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := `[{"Name":"Bob","Age":32,"Company":"IBM","Engineer":true},{"Name":"John","Age":20,"Company":"Oracle","Engineer":false},{"Name":"Henry","Age":45,"Company":"Microsoft","Engineer":false}]`
	for i:=0; i< t.N; i++ {
		_,_ = toolConvert.JsonToMap(str)
	}
}

///////////////////////////////////// 测试 ToInt ///////////////////////////////////
// 测试命令: go test -v -run TestToInt Convert/*
func TestToInt(t *testing.T) {
	var toolConvert Convert
	type Investment struct {
		Price  float64
		Symbol string
		Rating int64
	}
	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}
	tests := []struct {
		input  interface{}
	}{
		{int(9223372036854775807)},
		{int8(127)},
		{int16(32767)},
		{int32(2147483647)},
		{int64(9223372036854775807)},
		{uint(18446744073709551615)},
		{uint8(255)},
		{uint16(65535)},
		{uint32(4294967295)},
		{uint64(18446744073709551615)},
		{float32(8.31)},
		{float64(8.31)},
		{true},
		{false},
		{"9223372036854775807"},
		{nil},
		// errors
		{"test"},
		{testing.T{}},
		{&inv.Price},
		{&inv.Symbol},
		{&inv.Rating},
	}

	for _, test := range tests {
		//errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		//fmt.Println(errmsg)
		v,err := toolConvert.ToInt(test.input)
		b := reflect.ValueOf(test.input)
		fmt.Println(test.input , " >>>>>>>>>>>> " , b.Kind() , " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
	}
}

// go test -v -run TestToInt -bench=BenchmarkToInt -count=5 Convert/*
func BenchmarkToInt(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	type Investment struct {
		Price  float64
		Symbol string
		Rating int64
	}
	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}
	tests := []struct {
		input  interface{}
	}{
		{int(9223372036854775807)},
		{int8(127)},
		{int16(32767)},
		{int32(2147483647)},
		{int64(9223372036854775807)},
		{uint(18446744073709551615)},
		{uint8(255)},
		{uint16(65535)},
		{uint32(4294967295)},
		{uint64(18446744073709551615)},
		{float32(8.31)},
		{float64(8.31)},
		{true},
		{false},
		{"9223372036854775807"},
		{nil},
		// errors
		{"test"},
		{testing.T{}},
		{&inv.Price},
		{&inv.Symbol},
		{&inv.Rating},
	}
	for i:=0; i< t.N; i++ {
		for _, test := range tests {
			_,_ = toolConvert.ToInt(test.input)
		}

	}
}

/////////////////////////////////////// 测试 ToInt8 ///////////////////////////////////
//// 测试命令: go test -v -run TestToInt8 Convert/*
//func TestToInt8(t *testing.T) {
//	var toolConvert Convert
//	type Investment struct {
//		Price  float64
//		Symbol string
//		Rating int64
//	}
//	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{"9223372036854775807"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//		{&inv.Price},
//		{&inv.Symbol},
//		{&inv.Rating},
//	}
//
//	for _, test := range tests {
//		//errmsg := fmt.Sprintf("i = %d", i) // assert helper message
//		//fmt.Println(errmsg)
//		v,err := toolConvert.ToInt8(test.input)
//		b := reflect.ValueOf(test.input)
//		fmt.Println(test.input , " >>>>>>>>>>>> " , b.Kind() , " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
//	}
//
//}
//
//// go test -v -run TestToInt8 -bench=BenchmarkToInt8 -count=5 Convert/*
//func BenchmarkToInt8(t *testing.B) {
//	t.ResetTimer()
//	var toolConvert Convert
//	type Investment struct {
//		Price  float64
//		Symbol string
//		Rating int64
//	}
//	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{"9223372036854775807"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//		{&inv.Price},
//		{&inv.Symbol},
//		{&inv.Rating},
//	}
//	for i:=0; i< t.N; i++ {
//		for _, test := range tests {
//			_,_ = toolConvert.ToInt8(test.input)
//		}
//
//	}
//}

/////////////////////////////////////// 测试 ToInt16 ///////////////////////////////////
//// 测试命令: go test -v -run TestToInt16 Convert/*
//func TestToInt16(t *testing.T) {
//	var toolConvert Convert
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(88888888.31)},
//		{float64(88888888.31)},
//		{true},
//		{false},
//		{"9223372036854775807"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//
//	for _, test := range tests {
//		//errmsg := fmt.Sprintf("i = %d", i) // assert helper message
//		//fmt.Println(errmsg)
//		v,err := toolConvert.ToInt16(test.input)
//		b := reflect.ValueOf(test.input)
//		fmt.Println(test.input , " >>>>>>>>>>>> " , b.Kind() , " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
//	}
//}
//
//// go test -v -run TestToInt16 -bench=BenchmarkToInt16 -count=5 Convert/*
//func BenchmarkToInt16(t *testing.B) {
//	t.ResetTimer()
//	var toolConvert Convert
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{"8"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//	for i:=0; i< t.N; i++ {
//		for _, test := range tests {
//			_,_ = toolConvert.ToInt16(test.input)
//		}
//
//	}
//}
//
/////////////////////////////////////// 测试 ToInt32 ///////////////////////////////////
//// 测试命令: go test -v -run TestToInt32 Convert/*
//func TestToInt32(t *testing.T) {
//	var toolConvert Convert
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(88888888.31)},
//		{float64(88888888.31)},
//		{true},
//		{false},
//		{"9223372036854775807"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//
//	for _, test := range tests {
//		//errmsg := fmt.Sprintf("i = %d", i) // assert helper message
//		//fmt.Println(errmsg)
//		v,err := toolConvert.ToInt32(test.input)
//		b := reflect.ValueOf(test.input)
//		fmt.Println(test.input , " >>>>>>>>>>>> " , b.Kind() , " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
//	}
//}
//
//// go test -v -run TestToInt32 -bench=BenchmarkToInt32 -count=5 Convert/*
//func BenchmarkToInt32(t *testing.B) {
//	t.ResetTimer()
//	var toolConvert Convert
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{"8"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//	for i:=0; i< t.N; i++ {
//		for _, test := range tests {
//			_,_ = toolConvert.ToInt32(test.input)
//		}
//
//	}
//}
//
/////////////////////////////////////// 测试 ToInt64 ///////////////////////////////////
//// 测试命令: go test -v -run TestToInt64 Convert/*
//func TestToInt64(t *testing.T) {
//	var toolConvert Convert
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(88888888.31)},
//		{float64(88888888.31)},
//		{true},
//		{false},
//		{"9223372036854775807"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//
//	for _, test := range tests {
//		//errmsg := fmt.Sprintf("i = %d", i) // assert helper message
//		//fmt.Println(errmsg)
//		v,err := toolConvert.ToInt64(test.input)
//		b := reflect.ValueOf(test.input)
//		fmt.Println(test.input , " >>>>>>>>>>>> " , b.Kind() , " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
//	}
//}
//
//// go test -v -run TestToInt64 -bench=BenchmarkToInt64 -count=5 Convert/*
//func BenchmarkToInt64(t *testing.B) {
//	t.ResetTimer()
//	var toolConvert Convert
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{"8"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//	for i:=0; i< t.N; i++ {
//		for _, test := range tests {
//			_,_ = toolConvert.ToInt64(test.input)
//		}
//
//	}
//}