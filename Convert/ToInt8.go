package Convert

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// ToInt8 强制将变量转换为int8类型.
func (*Convert)ToInt8(value interface{}) (int8, error) {
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
		if s >= math.MinInt8 && s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case int8:
		return s, nil
	case int16:
		if s >= math.MinInt8 && s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case int32:
		if s >= math.MinInt8 && s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case int64:
		if s >= math.MinInt8 && s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case uint:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case uint8:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case uint16:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case uint32:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case uint64:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case float32:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case float64:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case string:
		vv, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int8(vv), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	}
}