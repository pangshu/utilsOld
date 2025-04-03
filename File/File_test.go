package File

import (
    "fmt"
    "testing"
)

///////////////////////////////////// 测试 AbsPath ///////////////////////////////////
// 测试命令: go test -v -run TestAbsPath File/*
func TestAbsPath(t *testing.T) {
    var toolFile File
    str := "./test.txt"
    res1 := toolFile.AbsPath(str)
    fmt.Println(res1)
}

// go test -v -run TestAbsPath -bench=BenchmarkAbsPath -count=5 File/*
func BenchmarkAbsPath(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    str := "./test.txt"
    for i:=0; i< t.N; i++ {
        _ = toolFile.AbsPath(str)
    }
}

///////////////////////////////////// 测试 AppendFile ///////////////////////////////////
// 测试命令: go test -v -run TestAppendFile File/*
func TestAppendFile(t *testing.T) {
    var toolFile File
    filename := "./test.txt"
    str := "测试测试"
    res1 := toolFile.AppendFile(filename, str)
    fmt.Println(res1)
}

// go test -v -run TestAppendFile -bench=BenchmarkAppendFile -count=5 File/*
func BenchmarkAppendFile(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    filename := "./test.txt"
    str := "测试测试"
    for i:=0; i< t.N; i++ {
        _ = toolFile.AppendFile(filename, str)
    }
}

///////////////////////////////////// 测试 Basename ///////////////////////////////////
// 测试命令: go test -v -run TestBasename File/*
func TestBasename(t *testing.T) {
    var toolFile File
    filename := "./test.txt"
    res1 := toolFile.Basename(filename)
    fmt.Println(res1)
}

// go test -v -run TestBasename -bench=BenchmarkBasename -count=5 File/*
func BenchmarkBasename(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    filename := "./test.txt"
    for i:=0; i< t.N; i++ {
        _ = toolFile.Basename(filename)
    }
}

///////////////////////////////////// 测试 Chmod ///////////////////////////////////
// 测试命令: go test -v -run TestChmod File/*
func TestChmod(t *testing.T) {
    var toolFile File
    filename := "./test.txt"
    res1 := toolFile.Chmod(filename, 0777)
    fmt.Println(res1)
}

// go test -v -run TestChmod -bench=BenchmarkChmod -count=5 File/*
func BenchmarkChmod(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    filename := "./test.txt"
    for i:=0; i< t.N; i++ {
        _ = toolFile.Chmod(filename, 0777)
    }
}

///////////////////////////////////// 测试 CopyDir ///////////////////////////////////
// 测试命令: go test -v -run TestCopyDir File/*
func TestCopyDir(t *testing.T) {
    var toolFile File
    src := "/www/golang/src/v4/aaa"
    dst := "/www/golang/src/v4/ccc"
    res1,_ := toolFile.CopyDir(src, dst)
    fmt.Println(res1)
}

// go test -v -run TestChmod -bench=BenchmarkCopyDir -count=5 File/*
func BenchmarkCopyDir(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    src := "./aaa"
    dst := "./bbb"
    for i:=0; i< t.N; i++ {
        _,_ = toolFile.CopyDir(src, dst)
    }
}

///////////////////////////////////// 测试 CopyFile ///////////////////////////////////
// 测试命令: go test -v -run TestCopyFile File/*
func TestCopyFile(t *testing.T) {
    var toolFile File
    src := "/www/golang/src/v4/bbb/test.txt"
    dst := "/www/golang/src/v4/bbb/bbb.txt"
    res1,_ := toolFile.CopyFile(src, dst)
    fmt.Println(res1)
}

// go test -v -run TestChmod -bench=BenchmarkFile -count=5 File/*
func BenchmarkCopyFile(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    src := "./aaa"
    dst := "./bbb"
    for i:=0; i< t.N; i++ {
        _,_ = toolFile.CopyFile(src, dst)
    }
}

///////////////////////////////////// 测试 CopyLink ///////////////////////////////////
// 测试命令: go test -v -run TestCopyLink File/*
func TestCopyLink(t *testing.T) {
    var toolFile File
    src := "/www/golang/src/v4/bbb/aaa.txt"
    dst := "/www/golang/src/v4/bbb/ccc.txt"
    res1 := toolFile.CopyLink(src, dst)
    fmt.Println(res1)
}

// go test -v -run TestChmod -bench=BenchmarkCopyLink -count=5 File/*
func BenchmarkCopyLink(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    src := "./aaa"
    dst := "./bbb"
    for i:=0; i< t.N; i++ {
        _ = toolFile.CopyLink(src, dst)
    }
}

///////////////////////////////////// 测试 CountLines ///////////////////////////////////
// 测试命令: go test -v -run TestCountLines File/*
func TestCountLines(t *testing.T) {
    var toolFile File
    src := "/www/golang/src/v4/bbb/aaa.txt"
    //dst := "/www/golang/src/v4/bbb/ccc.txt"
    res1,_ := toolFile.CountLines(src)
    fmt.Println(res1)
}

// go test -v -run TestChmod -bench=BenchmarkCountLines -count=5 File/*
func BenchmarkCountLines(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    src := "./aaa"
    for i:=0; i< t.N; i++ {
        _,_ = toolFile.CountLines(src)
    }
}

///////////////////////////////////// 测试 DelDir ///////////////////////////////////
// 测试命令: go test -v -run TestDelDir File/*
func TestDelDir(t *testing.T) {
    var toolFile File
    src := "/www/golang/src/v4/aaa"
    //dst := "/www/golang/src/v4/bbb/ccc.txt"
    res1 := toolFile.DelDir(src)
    fmt.Println(res1)
}

// go test -v -run TestChmod -bench=BenchmarkDelDir -count=5 File/*
func BenchmarkDelDir(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    src := "./aaa"
    for i:=0; i< t.N; i++ {
        _ = toolFile.DelDir(src)
    }
}

///////////////////////////////////// 测试 Dirname ///////////////////////////////////
// 测试命令: go test -v -run TestDirname File/*
func TestDirname(t *testing.T) {
    var toolFile File
    form := "/www/golang/src/v4/bbb"
    res1 := toolFile.Dirname(form)
    fmt.Println(res1)
}

// go test -v -run TestDirname -bench=BenchmarkDirname -count=5 File/*
func BenchmarkDirname(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    form := "/www/golang/src/v4/bbb/test.txt"
    for i:=0; i< t.N; i++ {
        _ = toolFile.Dirname(form)
    }
}

///////////////////////////////////// 测试 DirSize ///////////////////////////////////
// 测试命令: go test -v -run TestDirSize File/*
func TestDirSize(t *testing.T) {
    var toolFile File
    form := "/www/golang/src/v4/bbb/"
    res1 := toolFile.DirSize(form)
    fmt.Println(res1)
}

// go test -v -run TestDirSize -bench=BenchmarkDirSize -count=5 File/*
func BenchmarkDirSize(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    form := "/www/golang/src/v4/bbb/"
    for i:=0; i< t.N; i++ {
        _ = toolFile.DirSize(form)
    }
}

///////////////////////////////////// 测试 Ext ///////////////////////////////////
// 测试命令: go test -v -run TestExt File/*
func TestExt(t *testing.T) {
    var toolFile File
    form := "/www/golang/src/v4/bbb/test.txt"
    res1 := toolFile.Ext(form)
    fmt.Println(res1)
}

// go test -v -run TestExt -bench=BenchmarkExt -count=5 File/*
func BenchmarkExt(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    form := "/www/golang/src/v4/bbb/test"
    for i:=0; i< t.N; i++ {
        _ = toolFile.Ext(form)
    }
}

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////
// 测试命令: go test -v -run TestFastCopy File/*
func TestFastCopy(t *testing.T) {
    var toolFile File
    form := "/www/golang/src/v4/bbb/"
    to := "/www/golang/src/v4/ccc/"
    res1,_ := toolFile.FastCopy(form, to)
    fmt.Println(res1)
}

// go test -v -run TestFastCopy -bench=BenchmarkFastCopy -count=5 File/*
func BenchmarkFastCopy(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    form := "/www/golang/src/v4/bbb/"
    to := "/www/golang/src/v4/ccc/"
    for i:=0; i< t.N; i++ {
        _,_ = toolFile.FastCopy(form, to)
    }
}

///////////////////////////////////// 测试 FormatDir ///////////////////////////////////
// 测试命令: go test -v -run TestFormatDir File/*
func TestFormatDir(t *testing.T) {
    var toolFile File
    form := "/www/golang/src/v4/bbb/"
    res1 := toolFile.FormatDir(form)
    fmt.Println(res1)
}

// go test -v -run TestFormatDir -bench=BenchmarkFormatDir -count=5 File/*
func BenchmarkFormatDir(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    form := "/www/golang/src/v4/bbb/"
    for i:=0; i< t.N; i++ {
        _ = toolFile.FormatDir(form)
    }
}

///////////////////////////////////// 测试 FormatPath ///////////////////////////////////
// 测试命令: go test -v -run TestFormatPath File/*
func TestFormatPath(t *testing.T) {
    var toolFile File
    form := `/usr|///tmp:\\\123/\abc<|\hello>\/%world?\\how$\\are\@#test.png`
    res1 := toolFile.FormatPath(form)
    fmt.Println(res1)
}

// go test -v -run TestFormatPath -bench=BenchmarkFormatPath -count=5 File/*
func BenchmarkFormatPath(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    form := "/www/golang/src/v4"
    for i:=0; i< t.N; i++ {
        _ = toolFile.FormatPath(form)
    }
}

///////////////////////////////////// 测试 Glob ///////////////////////////////////
// 测试命令: go test -v -run TestGlob File/*
func TestGlob(t *testing.T) {
    var toolFile File
    form := "*test.go"
    res1,_ := toolFile.Glob(form)
    fmt.Println(res1)
}

// go test -v -run TestGlob -bench=BenchmarkGlob -count=5 File/*
func BenchmarkGlob(t *testing.B) {
    t.ResetTimer()
    var toolFile File
    form := "/www/golang/src/v4"
    for i:=0; i< t.N; i++ {
        _,_ = toolFile.Glob(form)
    }
}

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

///////////////////////////////////// 测试 FastCopy ///////////////////////////////////

