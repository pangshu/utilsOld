package Convert

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// ToInt 强制将变量转换为整型.
func (*Convert)ToInt(value interface{}) (int,error) {
	if value == nil {
		return 0, nil
	}

	i := indirect(value)
	v := reflect.ValueOf(i)
	//处理类型转换
	switch v.Kind() {
	case reflect.Invalid:
		return 0, nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		} else {
			return 0, nil
		}
	case reflect.String:
		vv, err := strconv.ParseInt(v.String(), 0, 0)
		if err == nil {
			return int(vv), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
		}
	case reflect.Int:
		return int(v.Int()), nil
	case reflect.Int8:
		return int(v.Int()), nil
	case reflect.Int16:
		return int(v.Int()), nil
	case reflect.Int32:
		return int(v.Int()), nil
	case reflect.Int64:
		if v.Int() >= math.MinInt && v.Int() <= math.MaxInt {
			return int(v.Int()), nil
		} else {
			return int(v.Int()), fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
		}
	case reflect.Uint:
		if v.Uint() <= math.MaxInt {
			return int(v.Uint()), nil
		} else {
			return int(v.Uint()), fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
		}
	case reflect.Uint8:
		return int(v.Uint()), nil
	case reflect.Uint16:
		return int(v.Uint()), nil
	case reflect.Uint32:
		intSize := 32 << (^uint(0) >> 63) // 32 or 64
		if math.MaxUint32 >= (1<<(intSize-1) - 1) {
			if v.Uint() <= (1<<(intSize-1) - 1) {
				return int(v.Uint()), nil
			} else {
				return int(v.Uint()), fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
			}
		} else {
			return int(v.Uint()), nil
		}
	case reflect.Uint64:
		if v.Uint() <= math.MaxInt {
			return int(v.Uint()), nil
		} else {
			return int(v.Uint()), fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
		}
	case reflect.Uintptr:
		if v.Uint() <= math.MaxInt {
			return int(v.Uint()), nil
		} else {
			return int(v.Uint()), fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
		}
	case reflect.Float32:
		if v.Float() <= math.MaxInt {
			return int(v.Float()), nil
		} else {
			return int(v.Float()), fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
		}
	case reflect.Float64:
		if v.Float() <= math.MaxInt {
			return int(v.Float()), nil
		} else {
			return int(v.Float()), fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
		}
	case reflect.Slice:
		vv, err := strconv.ParseInt(string(v.Bytes()), 0, 0)
		if err == nil {
			return int(vv), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
		}
	case reflect.Ptr, reflect.Struct, reflect.Map: //指针、结构体和字典
		return 0, fmt.Errorf("unable to cast %#v of type %v to string", v, v.Kind())
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %v to string", v, v.Kind())
	}




	//var i interface{}
	//if t := reflect.TypeOf(value); t.Kind() != reflect.Ptr {
	//	i = value
	//} else {
	//	v := reflect.ValueOf(value)
	//	for v.Kind() == reflect.Ptr && !v.IsNil() {
	//		v = v.Elem()
	//	}
	//	i = v.Interface()
	//}
	//
	////处理类型转换
	//switch s := i.(type) {
	//case int:
	//	return s, nil
	//case int8:
	//	return int(s), nil
	//case int16:
	//	return int(s), nil
	//case int32:
	//	return int(s), nil
	//case int64:
	//	if s >= math.MinInt && s <= math.MaxInt {
	//		return int(s), nil
	//	} else {
	//		return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	//	}
	//case uint:
	//	if s <= math.MaxInt {
	//		return int(s), nil
	//	} else {
	//		return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	//	}
	//case uint8:
	//	return int(s), nil
	//case uint16:
	//	return int(s), nil
	//case uint32:
	//	intSize := 32 << (^uint(0) >> 63) // 32 or 64
	//	if math.MaxUint32 >= (1<<(intSize-1) - 1) {
	//		if s <= (1<<(intSize-1) - 1) {
	//			return int(s), nil
	//		} else {
	//			return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	//		}
	//	} else {
	//		return int(s), nil
	//	}
	//case uint64:
	//	if s <= math.MaxInt {
	//		return int(s), nil
	//	} else {
	//		return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	//	}
	//case float32:
	//	if s <= math.MaxInt {
	//		return int(s), nil
	//	} else {
	//		return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	//	}
	//case float64:
	//	if s <= math.MaxInt {
	//		return int(s), nil
	//	} else {
	//		return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	//	}
	//case string:
	//	vv, err := strconv.ParseInt(s, 0, 0)
	//	if err == nil {
	//		return int(vv), nil
	//	} else {
	//		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	//	}
	//case bool:
	//	if s {
	//		return 1, nil
	//	} else {
	//		return 0, nil
	//	}
	//case nil:
	//	return 0, nil
	//default:
	//	return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	//}
}