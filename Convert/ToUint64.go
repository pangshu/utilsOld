package Convert

import (
	"fmt"
	"math"
	"strconv"
)

// ToUint64 强制将变量转换为int64类型.
func (*Convert)ToUint64(value interface{}) (uint64, error) {
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
	//	vv, err := strconv.ParseInt(v.String(), 8, 0)
	//	if err == nil {
	//		return uint64(vv), nil
	//	} else {
	//		return 0, fmt.Errorf("unable to cast %#v of type %v to uint64", v, v.Kind())
	//	}
	//case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	//	tmpValue := v.Int()
	//	if tmpValue >= 0 {
	//		return uint64(tmpValue), nil
	//	} else {
	//		return uint64(tmpValue), fmt.Errorf("unable to cast %#v of type %v to uint64", v, v.Kind())
	//	}
	//case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	//	tmpValue := v.Uint()
	//	if tmpValue <= math.MaxUint32 {
	//		return uint64(tmpValue), nil
	//	} else {
	//		return uint64(tmpValue), fmt.Errorf("unable to cast %#v of type %v to uint64", v, v.Kind())
	//	}
	//case reflect.Float32, reflect.Float64:
	//	tmpValue := v.Float()
	//	if tmpValue >= 0 && tmpValue <= math.MaxUint32 {
	//		return uint64(tmpValue), nil
	//	} else {
	//		return uint64(tmpValue), fmt.Errorf("unable to cast %#v of type %v to uint64", v, v.Kind())
	//	}
	//case reflect.Slice:
	//	vv, err := strconv.ParseInt(string(v.Bytes()), 8, 0)
	//	if err == nil {
	//		return uint64(vv), nil
	//	} else {
	//		return 0, fmt.Errorf("unable to cast %#v of type %v to uint64", v, v.Kind())
	//	}
	//case reflect.Ptr, reflect.Struct, reflect.Map: //指针、结构体和字典
	//	return 0, fmt.Errorf("unable to cast %#v of type %v to uint64", v, v.Kind())
	//default:
	//	return 0, fmt.Errorf("unable to cast %#v of type %v to uint64", v, v.Kind())
	//}



	//处理类型转换
	switch s := i.(type) {
	case int:
		if s >= 0 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case int8:
		if s >= 0 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case int16:
		if s >= 0 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case int32:
		if s >= 0 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case int64:
		if s >= 0 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case uint:
		return uint64(s), nil
	case uint8:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint32:
		return uint64(s), nil
	case uint64:
		return s, nil
	case float32:
		if s <= math.MaxUint64 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case float64:
		if s <= math.MaxUint64 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case string:
		vv, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return uint64(vv), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
	}
}
