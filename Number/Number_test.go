package Number

import (
	"fmt"
	"testing"
)

///////////////////////////////////// 测试 Average ///////////////////////////////////
// 测试命令: go test -v -run TestAverage Number/*
func TestAverage(t *testing.T) {
	var toolNumber Number
	num := []interface{}{1,2.2,3.3,4,5,6,7,8,9,10,"aa","bb"}
	res1 := toolNumber.Average(num)
	fmt.Println(res1)
}

// 测试命令: go test -v -run TestAverage -bench=BenchmarkAverage -count=5 Number/*
func BenchmarkAverage(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []interface{}{1,2,3,4,5,6,7,8,9,10}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Average(num)
	}
}
