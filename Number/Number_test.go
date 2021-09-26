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

///////////////////////////////////// 测试 Format ///////////////////////////////////
// 测试命令: go test -v -run TestFormat Number/*
func TestFormat(t *testing.T) {
	var toolNumber Number
	sum := 1234567890.123123123
	res1 := toolNumber.Format(sum, 8, ".", ",")
	fmt.Println(res1)
}

// go test -v -run TestFormat -bench=BenchmarkFormat -count=5 Number/*
func BenchmarkFormat(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	sum := 1234567890.123123123
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Format(sum, 2, ".", "")
	}
}

///////////////////////////////////// 测试 Sum ///////////////////////////////////
// 测试命令: go test -v -run TestSum Number/*
func TestSum(t *testing.T) {
	var toolNumber Number
	num := []interface{}{1,2,3,4,5,6,7,8,9,10,"aa","bb",[]string{"a","b"}}
	res1 := toolNumber.Sum(num)
	fmt.Println(res1)
}

// go test -v -run TestSum -bench=BenchmarkSum -count=5 Number/*
func BenchmarkSum(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []interface{}{1,2,3,4,5,6,7,8,9,10}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Sum(num)
	}
}