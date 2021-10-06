package Convert

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// ToUint32 强制将变量转换为int32类型.
func (*Convert)ToUint32(value interface{}) (uint32, error) {
	if value == nil {
		return 0, nil
	}
	var i interface{}
	if t := reflect.TypeOf(value); t.Kind() != reflect.Ptr {
		i = value
	} else {
		v := reflect.ValueOf(value)
		for v.Kind() == reflect.Ptr && !v.IsNil() {
			v = v.Elem()
		}
		i = v.Interface()
	}
	//处理类型转换
	switch s := i.(type) {
	case int:
		if s >= 0 && s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case int8:
		if s >= 0 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case int16:
		if s >= 0 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case int32:
		if s >= 0 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case int64:
		if s >= 0 && s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case uint:
		if s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case uint8:
		return uint32(s), nil
	case uint16:
		return uint32(s), nil
	case uint32:
		return uint32(s), nil
	case uint64:
		if s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case float32:
		if s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case float64:
		if s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case string:
		vv, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return uint32(vv), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
	}
}
