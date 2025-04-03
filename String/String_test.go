package String

import (
	"fmt"
	"testing"
)


///////////////////////////////////// 测试 ToCharset ///////////////////////////////////
// 测试命令: go test -v -run TestFirstLetter String/*
func TestFirstLetter(t *testing.T) {
	var toolString String
	res1 := toolString.FirstLetter("测Hello world. I love Shanghai!")
	fmt.Println(res1)
}

// go test -v -run TestFirstLetter -bench=BenchmarkFirstLetter -count=5 String/*
func BenchmarkFirstLetter(t *testing.B) {
	t.ResetTimer()
	for i:=0; i< t.N; i++ {
		var toolString String
		_ = toolString.FirstLetter("Hello world. I love Shanghai!")
	}
}

///////////////////////////////////// 测试 ToCharset ///////////////////////////////////
// 测试命令: go test -v -run TestExplode String/*
func TestExplode(t *testing.T) {
	var toolString String
	tmpStr := []string{" ", "(", "（"}
	res1 := toolString.Explode("测Hello", tmpStr...)
	//res1 := toolString.Explode("测Hello (world). I love Shanghai!", tmpStr...)
	fmt.Println(res1)
}
