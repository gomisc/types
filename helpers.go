package types

import "reflect"

// IsPrimitiveValues возвращает признак того что переданные типы являются примитивными
func IsPrimitiveValues(values ...interface{}) bool {
	for i := 0; i < len(values); i++ {
		v := reflect.TypeOf(values[i])
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64, reflect.Uintptr, reflect.String, reflect.Bool:
			continue
		default:
			return false
		}
	}

	return true
}
