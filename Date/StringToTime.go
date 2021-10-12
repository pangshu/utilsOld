package Date

import (
    "errors"
    "strings"
    "time"
)

// StringToTime 将字符串转换为时间结构.
// str 为要转换的字符串;
// format 为该字符串的格式,默认为"2006-01-02 15:04:05" .
func (*Date)StringToTime(str string, format ...string) (time.Time, error) {
    var f string
    if len(format) > 0 {
        f = strings.Trim(format[0], " ")
    } else {
        f = "2006-01-02 15:04:05"
    }

    if len(str) != len(f) {
        return time.Now(), errors.New("StrToTimestruct: parameter format error")
    }

    return time.ParseInLocation(f, str, time.Local)
}

