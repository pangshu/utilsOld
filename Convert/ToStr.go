package Convert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// ToStr 强制将变量转换为字符串.
func (*Convert)ToStr(value interface{}) (string,error) {
	if value == nil {
		return "", nil
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
		return "", nil
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil
	case reflect.String:
		return v.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32), nil
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), nil
	case reflect.Slice:
		return string(v.Bytes()), nil
	case reflect.Ptr, reflect.Struct, reflect.Map: //指针、结构体和字典
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return "", fmt.Errorf("unable to cast %#v of type %v to string", v, v)
		} else {
			return string(b), nil
		}
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", v, v)
	}

	////转换过程中，变量类型不能完全识别，只能手工后期增加
	//i := v.Interface()
	//switch s := i.(type) {
	//case string:
	//	return s, nil
	//case bool:
	//	return strconv.FormatBool(s), nil
	//case float64:
	//	return strconv.FormatFloat(s, 'f', -1, 64), nil
	//case float32:
	//	return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	//case int:
	//	return strconv.Itoa(s), nil
	//case int64:
	//	return strconv.FormatInt(s, 10), nil
	//case int32:
	//	return strconv.Itoa(int(s)), nil
	//case int16:
	//	return strconv.FormatInt(int64(s), 10), nil
	//case int8:
	//	return strconv.FormatInt(int64(s), 10), nil
	//case uint:
	//	return strconv.FormatUint(uint64(s), 10), nil
	//case uint64:
	//	return strconv.FormatUint(uint64(s), 10), nil
	//case uint32:
	//	return strconv.FormatUint(uint64(s), 10), nil
	//case uint16:
	//	return strconv.FormatUint(uint64(s), 10), nil
	//case uint8:
	//	return strconv.FormatUint(uint64(s), 10), nil
	//case []byte:
	//	return string(s), nil
	//case template.HTML:
	//	return string(s), nil
	//case template.URL:
	//	return string(s), nil
	//case template.JS:
	//	return string(s), nil
	//case template.CSS:
	//	return string(s), nil
	//case template.HTMLAttr:
	//	return string(s), nil
	//case nil:
	//	return "", nil
	//case fmt.Stringer:
	//	return s.String(), nil
	//case error:
	//	return s.Error(), nil
	//default:
	//	return "", fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	//}
}