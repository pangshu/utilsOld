package String

import (
	"fmt"
	"testing"
)


///////////////////////////////////// 测试 ToCharset ///////////////////////////////////
// 测试命令: go test -v -run TestFirstLetter String/*
func TestFirstLetter(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.FirstLetter("测Hello world. I love Shanghai!")
	fmt.Println(res1)
}

// go test -v -run TestFirstLetter -bench=BenchmarkFirstLetter -count=5 String/*
func BenchmarkFirstLetter(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.FirstLetter("Hello world. I love Shanghai!")
	}
}