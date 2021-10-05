package Convert

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToInt 强制将变量转换为整型.
func (*Convert)ToInt(value interface{}) int {
	if value == nil {
		return 0
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(value)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	//处理类型转换
	switch v.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Bool:
		if v.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.String:
		val, err := strconv.ParseInt(v.String(), 0, 0)
		if err == nil {
			return int(val)
		} else {
			return 0
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int(v.Uint())
	case reflect.Float32:
		return int(v.Float())
	case reflect.Float64:
		return int(v.Float())
	case reflect.Slice:
		return 0
	case reflect.Ptr, reflect.Struct, reflect.Map: //指针、结构体和字典
		return 0
	default:
		return 0
	}
	//
	//switch s := v.(type) {
	//case int:
	//	return s
	//case int64:
	//	return int(s)
	//case int32:
	//	return int(s)
	//case int16:
	//	return int(s)
	//case int8:
	//	return int(s)
	//case uint:
	//	return int(s)
	//case uint64:
	//	return int(s)
	//case uint32:
	//	return int(s)
	//case uint16:
	//	return int(s)
	//case uint8:
	//	return int(s)
	//case float64:
	//	return int(s)
	//case float32:
	//	return int(s)
	//case string:
	//	v, err := strconv.ParseInt(s, 0, 0)
	//	if err == nil {
	//		return int(v)
	//	}
	//	return 0
	//case bool:
	//	if s {
	//		return 1
	//	}
	//	return 0
	//case nil:
	//	return 0
	//default:
	//	return 0
	//}
}