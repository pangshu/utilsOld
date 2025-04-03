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

///////////////////////////////////// 测试 IsBetween ///////////////////////////////////
// 测试命令: go test -v -run TestIsBetween Number/*
func TestIsBetween(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.IsBetween(2,"1.9",5)
	fmt.Println(res1)
}

// 测试命令: go test -v -run TestIsBetween -bench=BenchmarkIsBetween -count=5 Number/*
func BenchmarkIsBetween(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.IsBetween(2,1,3)
	}
}

///////////////////////////////////// 测试 IsWhole ///////////////////////////////////
// 测试命令: go test -v -run TestIsWhole Number/*
func TestIsWhole(t *testing.T) {
	var toolNumber Number
	num := 5.0
	res1 := toolNumber.IsWhole(num)
	fmt.Println(res1)
}

// 测试命令: go test -v -run TestIsWhole -bench=BenchmarkIsWhole -count=5 Number/*
func BenchmarkIsWhole(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := 5.5
	for i:=0; i< t.N; i++ {
		_ = toolNumber.IsWhole(num)
	}
}

///////////////////////////////////// 测试 Max ///////////////////////////////////
// 测试命令: go test -v -run TestMax Number/*
func TestMax(t *testing.T) {
	var toolNumber Number
	num := []interface{}{-1, 0, "18", true, nil, int8(1), int16(2), int32(3), int64(4), uint(5),
		uint8(6), uint16(7), uint32(8), uint64(9), float32(10.0), float64(11.1)}
	res1 := toolNumber.Max(num)
	fmt.Println(res1)
}

// 测试命令: go test -v -run TestMax -bench=BenchmarkMax -count=5 Number/*
func BenchmarkMax(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []interface{}{-1, 0, "18", true, nil, int8(1), int16(2), int32(3), int64(4), uint(5),
		uint8(6), uint16(7), uint32(8), uint64(9), float32(10.0), float64(11.1)}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Max(num)
	}
}

///////////////////////////////////// 测试 Min ///////////////////////////////////
// 测试命令: go test -v -run TestMin Number/*
func TestMin(t *testing.T) {
	var toolNumber Number
	num := []interface{}{-1, 0, "18", true, nil, int8(1), int16(2), int32(3), int64(4), uint(5),
		uint8(6), uint16(7), uint32(8), uint64(9), float32(10.0), float64(11.1)}
	res1 := toolNumber.Min(num)
	fmt.Println(res1)
}

// 测试命令: go test -v -run TestMin -bench=BenchmarkMin -count=5 Number/*
func BenchmarkMin(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []interface{}{-1, 0, "18", true, nil, int8(1), int16(2), int32(3), int64(4), uint(5),
		uint8(6), uint16(7), uint32(8), uint64(9), float32(10.0), float64(11.1)}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Min(num)
	}
}

///////////////////////////////////// 测试 Percent ///////////////////////////////////
// 测试命令: go test -v -run TestPercent Number/*
func TestPercent(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.Percent(1, 1000)
	fmt.Println(res1)
}

// 测试命令: go test -v -run TestPercent -bench=BenchmarkPercent -count=5 Number/*
func BenchmarkPercent(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Percent(5, 10)
	}
}

///////////////////////////////////// 测试 RandArray ///////////////////////////////////
// 测试命令: go test -v -run TestRandArray Number/*
func TestRandArray(t *testing.T) {
	var toolNumber Number
	res1,resErr := toolNumber.RandArray(1,6, 6, false)
	fmt.Println(res1)
	fmt.Println(resErr)
}

// 测试命令: go test -v -run TestRandArray -bench=BenchmarkRandArray -count=5 Number/*
func BenchmarkRandArray(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_,_ = toolNumber.RandArray(5,10, 3, false)
	}
}

///////////////////////////////////// 测试 RandInt ///////////////////////////////////
// 测试命令: go test -v -run TestRandInt Number/*
func TestRandInt(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.RandInt(100,1)
	fmt.Println(res1)
}

// 测试命令: go test -v -run TestRandInt -bench=BenchmarkRandInt -count=5 Number/*
func BenchmarkRandInt(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.RandInt(1,10)
	}
}

///////////////////////////////////// 测试 RandInt64 ///////////////////////////////////
// 测试命令: go test -v -run TestRandInt64 Number/*
func TestRandInt64(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.RandInt64(100,1)
	fmt.Println(res1)
}

// 测试命令: go test -v -run TestRandInt64 -bench=BenchmarkRandInt64 -count=5 Number/*
func BenchmarkRandInt64(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.RandInt64(1,10)
	}
}

///////////////////////////////////// 测试 Round ///////////////////////////////////
// 测试命令: go test -v -run TestRound Number/*
func TestRound(t *testing.T) {
	var toolNumber Number
	num := 123.81515151515
	res1 := toolNumber.Round(num,3)
	fmt.Println(res1)
}

// 测试命令: go test -v -run TestRound -bench=BenchmarkRound -count=5 Number/*
func BenchmarkRound(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := 123.815
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Round(num,3)
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