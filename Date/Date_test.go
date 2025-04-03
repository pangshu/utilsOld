package Date

import (
    "fmt"
    "testing"
    "time"
)

///////////////////////////////////// 测试 CountDays ///////////////////////////////////
// 测试命令: go test -v -run TestCountDays Date/*
func TestCountDays(t *testing.T) {
    var toolDate Date
    res1 := toolDate.CountDays(time.Unix(1610236800,0), time.Now())
    fmt.Println(res1)
}

// go test -v -run TestCountDays -bench=BenchmarkCountDays -count=5 Date/*
func BenchmarkCountDays(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.CountDays(time.Unix(1610236800,0), time.Now())
    }
}

///////////////////////////////////// 测试 Day ///////////////////////////////////
// 测试命令: go test -v -run TestDay Date/*
func TestDay(t *testing.T) {
    var toolDate Date
    res1 := toolDate.Day(time.Now())
    fmt.Println(res1)
}

// go test -v -run TestDay -bench=BenchmarkDay -count=5 Date/*
func BenchmarkDay(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.Day(time.Now())
    }
}

///////////////////////////////////// 测试 EndOfDay ///////////////////////////////////
// 测试命令: go test -v -run TestEndOfDay Date/*
func TestEndOfDay(t *testing.T) {
    var toolDate Date
    res1 := toolDate.EndOfDay(time.Now())
    fmt.Println(res1.Unix())
}

// go test -v -run TestEndOfDay -bench=BenchmarkEndOfDay -count=5 Date/*
func BenchmarkEndOfDay(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.EndOfDay(time.Now())
    }
}

///////////////////////////////////// 测试 EndOfMonth ///////////////////////////////////
// 测试命令: go test -v -run TestEndOfMonth Date/*
func TestEndOfMonth(t *testing.T) {
    var toolDate Date
    res1 := toolDate.EndOfMonth(time.Now())
    fmt.Println(res1.Unix())
}

// go test -v -run TestEndOfMonth -bench=BenchmarkEndOfMonth -count=5 Date/*
func BenchmarkEndOfMonth(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.EndOfMonth(time.Now())
    }
}

///////////////////////////////////// 测试 EndOfWeek ///////////////////////////////////
// 测试命令: go test -v -run TestEndOfWeek Date/*
func TestEndOfWeek(t *testing.T) {
    var toolDate Date
    res1 := toolDate.EndOfWeek(time.Now())
    fmt.Println(res1.Unix())
}

// go test -v -run TestEndOfWeek -bench=BenchmarkEndOfWeek -count=5 Date/*
func BenchmarkEndOfWeek(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.EndOfWeek(time.Now())
    }
}

///////////////////////////////////// 测试 EndOfYear ///////////////////////////////////
// 测试命令: go test -v -run TestEndOfYear Date/*
func TestEndOfYear(t *testing.T) {
    var toolDate Date
    res1 := toolDate.EndOfYear(time.Now())
    fmt.Println(res1.Unix())
}

// go test -v -run TestEndOfYear -bench=BenchmarkEndOfYear -count=5 Date/*
func BenchmarkEndOfYear(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.EndOfYear(time.Now())
    }
}

///////////////////////////////////// 测试 Format ///////////////////////////////////
// 测试命令: go test -v -run TestFormat Date/*
func TestFormat(t *testing.T) {
    var toolDate Date
    res1 := toolDate.Format("Y-m-d H:i:s", time.Now())
    fmt.Println(res1)
}

// go test -v -run TestFormat -bench=BenchmarkFormat -count=5 Date/*
func BenchmarkFormat(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.Format("Y-m-d H:i:s", time.Now())
    }
}

///////////////////////////////////// 测试 Hour ///////////////////////////////////
// 测试命令: go test -v -run TestHour Date/*
func TestHour(t *testing.T) {
    var toolDate Date
    res1 := toolDate.Hour(time.Now())
    fmt.Println(res1)
}

// go test -v -run TestHour -bench=BenchmarkHour -count=5 Date/*
func BenchmarkHour(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.Hour(time.Now())
    }
}

///////////////////////////////////// 测试 IsDate ///////////////////////////////////
// 测试命令: go test -v -run TestIsDate Date/*
func TestIsDate(t *testing.T) {
    var toolDate Date
    res1 := toolDate.IsDate(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
    fmt.Println(res1)
}

// go test -v -run TestIsDate -bench=BenchmarkIsDate -count=5 Date/*
func BenchmarkIsDate(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.IsDate(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
    }
}

///////////////////////////////////// 测试 IsDateByString ///////////////////////////////////
// 测试命令: go test -v -run TestIsDateByString Date/*
func TestIsDateByString(t *testing.T) {
    var toolDate Date
    res1 := toolDate.IsDateByString("2021-01-19")
    fmt.Println(res1)
}

// go test -v -run TestIsDateByString -bench=BenchmarkIsDateByString -count=5 Date/*
func BenchmarkIsDateToTime(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.IsDateByString("2021-01-19")
    }
}

///////////////////////////////////// 测试 MonthDays ///////////////////////////////////
// 测试命令: go test -v -run TestMonthDays Date/*
func TestMonthDays(t *testing.T) {
    var toolDate Date
    res1 := toolDate.MonthDays(2)
    fmt.Println(res1)
}

// go test -v -run TestMonthDays -bench=BenchmarkMonthDays -count=5 Date/*
func BenchmarkMonthDays(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.MonthDays(2)
    }
}

///////////////////////////////////// 测试 MicroTime ///////////////////////////////////
// 测试命令: go test -v -run TestMicroTime Date/*
func TestMicroTime(t *testing.T) {
    var toolDate Date
    res1 := toolDate.MicroTime()
    fmt.Println(res1)
}

// go test -v -run TestMicroTime -bench=BenchmarkMicroTime -count=5 Date/*
func BenchmarkMicroTime(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.MicroTime()
    }
}

///////////////////////////////////// 测试 MilliTime ///////////////////////////////////
// 测试命令: go test -v -run TestMilliTime Date/*
func TestMilliTime(t *testing.T) {
    var toolDate Date
    res1 := toolDate.MilliTime()
    fmt.Println(res1)
}

// go test -v -run TestMilliTime -bench=BenchmarkMilliTime -count=5 Date/*
func BenchmarkMilliTime(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.MilliTime()
    }
}

///////////////////////////////////// 测试 Minute ///////////////////////////////////
// 测试命令: go test -v -run TestMinute Date/*
func TestMinute(t *testing.T) {
    var toolDate Date
    res1 := toolDate.Minute()
    fmt.Println(res1)
}

// go test -v -run TestMinute -bench=BenchmarkMinute -count=5 Date/*
func BenchmarkMinute(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.Minute()
    }
}

///////////////////////////////////// 测试 Month ///////////////////////////////////
// 测试命令: go test -v -run TestMonth Date/*
func TestMonth(t *testing.T) {
    var toolDate Date
    res1 := toolDate.Month()
    fmt.Println(res1)
}

// go test -v -run TestMonth -bench=BenchmarkMonth -count=5 Date/*
func BenchmarkMonth(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.Month()
    }
}

///////////////////////////////////// 测试 Second ///////////////////////////////////
// 测试命令: go test -v -run TestSecond Date/*
func TestSecond(t *testing.T) {
    var toolDate Date
    res1 := toolDate.Second()
    fmt.Println(res1)
}

// go test -v -run TestSecond -bench=BenchmarkSecond -count=5 Date/*
func BenchmarkSecond(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.Second()
    }
}

///////////////////////////////////// 测试 StartOfDay ///////////////////////////////////
// 测试命令: go test -v -run TestStartOfDay Date/*
func TestStartOfDay(t *testing.T) {
    var toolDate Date
    res1 := toolDate.StartOfDay(time.Now())
    fmt.Println(res1.Unix())
}

// go test -v -run TestStartOfDay -bench=BenchmarkStartOfDay -count=5 Date/*
func BenchmarkStartOfDay(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.StartOfDay(time.Now())
    }
}

///////////////////////////////////// 测试 StartOfMonth ///////////////////////////////////
// 测试命令: go test -v -run TestStartOfMonth Date/*
func TestStartOfMonth(t *testing.T) {
    var toolDate Date
    res1 := toolDate.StartOfMonth(time.Now())
    fmt.Println(res1.Unix())
}

// go test -v -run TestStartOfMonth -bench=BenchmarkStartOfMonth -count=5 Date/*
func BenchmarkStartOfMonth(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.StartOfMonth(time.Now())
    }
}

///////////////////////////////////// 测试 StartOfWeek ///////////////////////////////////
// 测试命令: go test -v -run TestStartOfWeek Date/*
func TestStartOfWeek(t *testing.T) {
    var toolDate Date
    res1 := toolDate.StartOfWeek(time.Now())
    fmt.Println(res1.Unix())
}

// go test -v -run TestStartOfWeek -bench=BenchmarkStartOfWeek -count=5 Date/*
func BenchmarkStartOfWeek(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.StartOfWeek(time.Now())
    }
}

///////////////////////////////////// 测试 StartOfYear ///////////////////////////////////
// 测试命令: go test -v -run TestStartOfYear Date/*
func TestStartOfYear(t *testing.T) {
    var toolDate Date
    res1 := toolDate.StartOfYear(time.Now())
    fmt.Println(res1.Unix())
}

// go test -v -run TestStartOfYear -bench=BenchmarkStartOfYear -count=5 Date/*
func BenchmarkStartOfYear(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.StartOfYear(time.Now())
    }
}

///////////////////////////////////// 测试 StringToTime ///////////////////////////////////
// 测试命令: go test -v -run TestStringToTime Date/*
func TestStringToTime(t *testing.T) {
    var toolDate Date
    res1,_ := toolDate.StringToTime("2019-07-11 10:11:23")
    fmt.Println(res1)
}

// go test -v -run TestStringToTime -bench=BenchmarkStringToTime -count=5 Date/*
func BenchmarkStringToTime(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _,_ = toolDate.StringToTime("2019-07-11 10:11:23")
    }
}

///////////////////////////////////// 测试 UnixTime ///////////////////////////////////
// 测试命令: go test -v -run TestUnixTime Date/*
func TestUnixTime(t *testing.T) {
    var toolDate Date
    res1 := toolDate.UnixTime()
    fmt.Println(res1)
}

// go test -v -run TestUnixTime -bench=BenchmarkUnixTime -count=5 Date/*
func BenchmarkUnixTime(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.UnixTime()
    }
}

///////////////////////////////////// 测试 Year ///////////////////////////////////
// 测试命令: go test -v -run TestYear Date/*
func TestYear(t *testing.T) {
    var toolDate Date
    res1 := toolDate.Year()
    fmt.Println(res1)
}

// go test -v -run TestYear -bench=BenchmarkYear -count=5 Date/*
func BenchmarkYear(t *testing.B) {
    t.ResetTimer()
    var toolDate Date
    for i:=0; i< t.N; i++ {
        _ = toolDate.Year()
    }
}
