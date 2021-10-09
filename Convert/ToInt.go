package Convert

import (
	"fmt"
	"math"
	"strconv"
)

// ToInt 强制将变量转换为整型.
func (*Convert)ToInt(value interface{}) (int,error) {
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
	//		return int(vv), nil
	//	} else {
	//		return 0, fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
	//	}
	//case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	//	tmpValue := v.Int()
	//	if tmpValue >= math.MinInt && tmpValue <= math.MaxInt {
	//		return int(tmpValue), nil
	//	} else {
	//		return int(tmpValue), fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
	//	}
	//case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	//	tmpValue := v.Uint()
	//	if tmpValue <= math.MaxInt {
	//		return int(tmpValue), nil
	//	} else {
	//		return int(tmpValue), fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
	//	}
	//case reflect.Float32, reflect.Float64:
	//	tmpValue := v.Float()
	//	if tmpValue <= math.MaxInt {
	//		return int(tmpValue), nil
	//	} else {
	//		return int(tmpValue), fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
	//	}
	//case reflect.Slice:
	//	vv, err := strconv.ParseInt(string(v.Bytes()), 0, 0)
	//	if err == nil {
	//		return int(vv), nil
	//	} else {
	//		return 0, fmt.Errorf("unable to cast %#v of type %v to int", v, v.Kind())
	//	}
	////case reflect.Ptr, reflect.Struct, reflect.Map: //指针、结构体和字典
	////	return 0, fmt.Errorf("unable to cast %#v of type %v to string", v, v.Kind())
	//default:
	//	return 0, fmt.Errorf("unable to cast %#v of type %v to string", v, v.Kind())
	//}




	//处理类型转换
	switch s := i.(type) {
	case int:
		return s, nil
	case int8:
		return int(s), nil
	case int16:
		return int(s), nil
	case int32:
		return int(s), nil
	case int64:
		if s >= math.MinInt && s <= math.MaxInt {
			return int(s), nil
		} else {
			return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case uint:
		if s <= math.MaxInt {
			return int(s), nil
		} else {
			return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case uint8:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint32:
		intSize := 32 << (^uint(0) >> 63) // 32 or 64
		if math.MaxUint32 >= (1<<(intSize-1) - 1) {
			if s <= (1<<(intSize-1) - 1) {
				return int(s), nil
			} else {
				return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
			}
		} else {
			return int(s), nil
		}
	case uint64:
		if s <= math.MaxInt {
			return int(s), nil
		} else {
			return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case float32:
		if s <= math.MaxInt {
			return int(s), nil
		} else {
			return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case float64:
		if s <= math.MaxInt {
			return int(s), nil
		} else {
			return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		vv, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(vv), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}