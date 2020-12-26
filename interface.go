package isodd

import (
	"reflect"
)

// interface type to check whether is odd, only support integer type
// float type, string type. Others are invalid.
func Interface(i interface{}) (bool, error) {
	t := reflect.TypeOf(i).Kind()

	switch t {
	case reflect.Int:
		n, _ := i.(int)
		return Int(n), nil
	case reflect.Int8:
		n, _ := i.(int8)
		return Int8(n), nil
	case reflect.Int16:
		n, _ := i.(int16)
		return Int16(n), nil
	case reflect.Int32:
		n, _ := i.(int32)
		return Int32(n), nil
	case reflect.Int64:
		n, _ := i.(int64)
		return Int64(n), nil
	}

	switch t {
	case reflect.Uint:
		n, _ := i.(uint)
		return Uint(n), nil
	case reflect.Uint8:
		n, _ := i.(uint8)
		return Uint8(n), nil
	case reflect.Uint16:
		n, _ := i.(uint16)
		return Uint16(n), nil
	case reflect.Uint32:
		n, _ := i.(uint32)
		return Uint32(n), nil
	case reflect.Uint64:
		n, _ := i.(uint64)
		return Uint64(n), nil
	}

	switch t {
	case reflect.String:
		n, _ := i.(string)
		return String(n)
	case reflect.Float32:
		n, _ := i.(float32)
		return Float32(n), nil
	case reflect.Float64:
		n, _ := i.(float64)
		return Float64(n), nil
	}

	return false, ErrorInterfaceNoMatch
}
