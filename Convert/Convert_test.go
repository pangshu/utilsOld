package Convert

import (
	"fmt"
	"html/template"
	"reflect"
	"testing"
)


///////////////////////////////////// 测试 ToStr ///////////////////////////////////
// 测试命令: go test -v -run TestToStr Convert/*
func TestToStr(t *testing.T) {
	var toolConvert Convert
	//str := []byte{230,181,139,232,175,149,228,184,173,230,150,135}
	//aaa := fmt.Sprintf("%T", str)
	//fmt.Println(aaa)
	//res1 := toolConvert.ToStr(str)
	//fmt.Println(res1)

	type Key struct {
		k string
	}
	key := &Key{"foo"}

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
		// errors
		{testing.T{}},
		{key},
	}

	for _, test := range tests {
		//errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		//fmt.Println(errmsg)
		v := toolConvert.ToStr(test.input)
		b := reflect.ValueOf(test.input)
		fmt.Println(test.input , " >>>>>>>>>>>> " , b.Kind() , " >>>>>>>>>>> ", v)
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
		// errors
		{testing.T{}},
		{key},
	}
	for i:=0; i< t.N; i++ {
		for _, test := range tests {
			_ = toolConvert.ToStr(test.input)
		}

	}
}