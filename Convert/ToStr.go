package Convert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// ToStr 强制将变量转换为字符串.
func (*Convert)ToStr(value interface{}) string {
	if value == nil {
		return ""
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
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Slice:
		return string(v.Bytes())
	case reflect.Ptr, reflect.Struct, reflect.Map: //指针、结构体和字典
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return ""
		}
		return string(b)
	default:
		return ""
	}

	//转换过程中，变量类型不能完全识别，只能手工后期增加
	//switch s := value.(type) {
	//case string:
	//	return s
	//case bool:
	//	return strconv.FormatBool(s)
	//case float64:
	//	return strconv.FormatFloat(s, 'f', -1, 64)
	//case float32:
	//	return strconv.FormatFloat(float64(s), 'f', -1, 32)
	//case int:
	//	return strconv.Itoa(s)
	//case int64:
	//	return strconv.FormatInt(s, 10)
	//case int32:
	//	return strconv.Itoa(int(s))
	//case int16:
	//	return strconv.FormatInt(int64(s), 10)
	//case int8:
	//	return strconv.FormatInt(int64(s), 10)
	//case uint:
	//	return strconv.FormatUint(uint64(s), 10)
	//case uint64:
	//	return strconv.FormatUint(uint64(s), 10)
	//case uint32:
	//	return strconv.FormatUint(uint64(s), 10)
	//case uint16:
	//	return strconv.FormatUint(uint64(s), 10)
	//case uint8:
	//	return strconv.FormatUint(uint64(s), 10)
	//case []byte:
	//	return string(s)
	//case template.HTML:
	//	return string(s)
	//case template.URL:
	//	return string(s)
	//case template.JS:
	//	return string(s)
	//case template.CSS:
	//	return string(s)
	//case template.HTMLAttr:
	//	return string(s)
	//case nil:
	//	return ""
	//case fmt.Stringer:
	//	return s.String()
	//case error:
	//	return s.Error()
	//default:
	//	return ""
	//}
}