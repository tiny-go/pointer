package pointer

import "reflect"

// Unpoint returns actual value of the pointers
func Unpoint(v interface{}) interface{} {
	// return nil if nil
	if v == nil {
		return nil
	}
	rv := reflect.ValueOf(v)
	// check if pointer
	if rv.Kind() == reflect.Ptr {
		// if zero vlue or a nil pointer
		if rv.IsNil() {
			return nil
		}
		// get pointer value recursively
		return Unpoint(reflect.Indirect(rv).Interface())
	}
	// it is an actual value
	return v
}
