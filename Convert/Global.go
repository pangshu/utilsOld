package Convert

import "reflect"

type Convert byte

func indirect(a interface{}) interface{} {
    if a == nil {
        return nil
    }
    if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
        // Avoid creating a reflect.Value if it's not a pointer.
        return a
    }
    v := reflect.ValueOf(a)
    for v.Kind() == reflect.Ptr && !v.IsNil() {
        v = v.Elem()
    }
    return v.Interface()
}