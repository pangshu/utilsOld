package Convert

import (
	"fmt"
	"strconv"
)

// ToBool 强制将变量转换为整型.
func (*Convert)ToBool(value interface{}) (bool,error) {
	if value == nil {
		return false, nil
	}
	i := indirect(value)

	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		if i.(int) != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(i.(string))
	default:
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", i, i)
	}
}