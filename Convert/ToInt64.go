package Convert

import (
	"fmt"
	"math"
	"strconv"
)

// ToInt64 强制将变量转换为int64类型.
func (*Convert)ToInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, nil
	}
	i := indirect(value)

	////处理类型转换
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//case reflect.Invalid:
	//	return 0, nil
	//case reflect.Bool:
	//	if v.Bool() {
	//		return 1, nil
	//	} else {
	//		return 0, nil
	//	}
	//case reflect.String:
	//	vv, err := strconv.ParseInt(v.String(), 0, 0)
	//	if err == nil {
	//		return vv, nil
	//	} else {
	//		return 0, fmt.Errorf("unable to cast %#v of type %v to int64", v, v.Kind())
	//	}
	//case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	//	return v.Int(), nil
	//
	//case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	//	tmpValue := v.Uint()
	//	if tmpValue <= math.MaxInt64 {
	//		return int64(tmpValue), nil
	//	} else {
	//		return int64(tmpValue), fmt.Errorf("unable to cast %#v of type %v to int64", v, v.Kind())
	//	}
	//case reflect.Float32, reflect.Float64:
	//	tmpValue := v.Float()
	//	if tmpValue <= math.MaxInt64 {
	//		return int64(tmpValue), nil
	//	} else {
	//		return int64(tmpValue), fmt.Errorf("unable to cast %#v of type %v to int64", v, v.Kind())
	//	}
	//case reflect.Slice:
	//	vv, err := strconv.ParseInt(string(v.Bytes()), 32, 0)
	//	if err == nil {
	//		return int64(vv), nil
	//	} else {
	//		return 0, fmt.Errorf("unable to cast %#v of type %v to int64", v, v.Kind())
	//	}
	//case reflect.Ptr, reflect.Struct, reflect.Map: //指针、结构体和字典
	//	return 0, fmt.Errorf("unable to cast %#v of type %v to int64", v, v.Kind())
	//default:
	//	return 0, fmt.Errorf("unable to cast %#v of type %v to int64", v, v.Kind())
	//}
	
	
	
	
	//处理类型转换
	switch s := i.(type) {
	case int:
		if s >= math.MinInt64 && s <= math.MaxInt64 {
			return int64(s), nil
		} else {
			return int64(s), fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case int8:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int32:
		return int64(s), nil

	case int64:
		return s, nil
	case uint:
		if s <= math.MaxInt64 {
			return int64(s), nil
		} else {
			return int64(s), fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case uint8:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint64:
		if s <= math.MaxInt64 {
			return int64(s), nil
		} else {
			return int64(s), fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case float32:
		if s <= math.MaxInt64 {
			return int64(s), nil
		} else {
			return int64(s), fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case float64:
		if s <= math.MaxInt64 {
			return int64(s), nil
		} else {
			return int64(s), fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case string:
		vv, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int64(vv), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case bool:
		if s {
			return 1, nil
		} else {
			return 0, nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	}
}
