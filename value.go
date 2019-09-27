package pointer

import "time"

// BoolVal returns the value of the boolean pointer and "true" or "false" and "false"
// if the pointer is nil
func BoolVal(ptr *bool) (val bool, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// UintVal returns the value of the uint pointer and "true" or 0 and "false"
// if the pointer is nil
func UintVal(ptr *uint) (val uint, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Uint8Val returns the value of the uint8 pointer and "true" or 0 and "false"
// if the pointer is nil
func Uint8Val(ptr *uint8) (val uint8, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Uint16Val returns the value of the uint16 pointer and "true" or 0 and "false"
// if the pointer is nil
func Uint16Val(ptr *uint16) (val uint16, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Uint32Val returns the value of the uint32 pointer and "true" or 0 and "false"
// if the pointer is nil
func Uint32Val(ptr *uint32) (val uint32, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Uint64Val returns the value of the uint64 pointer and "true" or 0 and "false"
// if the pointer is nil
func Uint64Val(ptr *uint64) (val uint64, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// UintptrVal returns the value of the uintptr pointer and "true" or 0 and "false"
// if the pointer is nil
func UintptrVal(ptr *uintptr) (val uintptr, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// IntVal returns the value of the int pointer and "true" or 0 and "false"
// if the pointer is nil
func IntVal(ptr *int) (val int, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Int8Val returns the value of the int8 pointer and "true" or 0 and "false"
// if the pointer is nil
func Int8Val(ptr *int8) (val int8, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Int16Val returns the value of the int16 pointer and "true" or 0 and "false"
// if the pointer is nil
func Int16Val(ptr *int16) (val int16, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Int32Val returns the value of the int32 pointer and "true" or 0 and "false"
// if the pointer is nil
func Int32Val(ptr *int32) (val int32, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Int64Val returns the value of the int64 pointer and "true" or 0 and "false"
// if the pointer is nil
func Int64Val(ptr *int64) (val int64, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Float32Val returns the value of the float32 pointer and "true" or 0 and "false"
// if the pointer is nil
func Float32Val(ptr *float32) (val float32, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Float64Val returns the value of the float64 pointer and "true" or 0 and "false"
// if the pointer is nil
func Float64Val(ptr *float64) (val float64, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// RuneVal returns the value of the rune pointer and "true" or 0 and "false"
// if the pointer is nil
func RuneVal(ptr *rune) (val rune, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// ByteVal returns the value of the byte pointer and "true" or 0 and "false"
// if the pointer is nil
func ByteVal(ptr *byte) (val byte, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// StringVal returns the value of the string pointer and "true" or an empty string
//  and "false" if the pointer is nil
func StringVal(ptr *string) (val string, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Complex64Val returns the value of the complex64 pointer and "true" or 0 and "false"
// if the pointer is nil
func Complex64Val(ptr *complex64) (val complex64, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// Complex128Val returns the value of the complex128 pointer and "true" or 0 and "false"
// if the pointer is nil
func Complex128Val(ptr *complex128) (val complex128, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// InterfaceVal returns the value of the interface{} pointer and "true" or nil and "false"
// if the pointer is nil
func InterfaceVal(ptr *interface{}) (val interface{}, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// DurationVal returns the value of the time.Duration pointer and "true" or 0 and "false"
// if the pointer is nil
func DurationVal(ptr *time.Duration) (val time.Duration, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}

// TimeVal returns the value of the time.Time pointer and "true" or zero time.Time
// and "false" if the pointer is nil
func TimeVal(ptr *time.Time) (val time.Time, ok bool) {
	if ptr != nil {
		return *ptr, true
	}
	return
}
