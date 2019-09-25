package pointer

import "time"

// Bool creates a pointer to the provided boolean value
func Bool(v bool) *bool { return &v }

// Uint creates a pointer to the provided uint value
func Uint(v uint) *uint { return &v }

// Uint8 creates a pointer to the provided uint8 value
func Uint8(v uint8) *uint8 { return &v }

// Uint16 creates a pointer to the provided uint16 value
func Uint16(v uint16) *uint16 { return &v }

// Uint32 creates a pointer to the provided uint32 value
func Uint32(v uint32) *uint32 { return &v }

// Uint64 creates a pointer to the provided uint64 value
func Uint64(v uint64) *uint64 { return &v }

// Uintptr creates a pointer to the provided uintptr value
func Uintptr(v uintptr) *uintptr { return &v }

// Int creates a pointer to the provided int value
func Int(v int) *int { return &v }

// Int8 creates a pointer to the provided int8 value
func Int8(v int8) *int8 { return &v }

// Int16 creates a pointer to the provided int16 value
func Int16(v int16) *int16 { return &v }

// Int32 creates a pointer to the provided int32 value
func Int32(v int32) *int32 { return &v }

// Int64 creates a pointer to the provided int64 value
func Int64(v int64) *int64 { return &v }

// Float32 creates a pointer to the provided float32 value
func Float32(v float32) *float32 { return &v }

// Float64 creates a pointer to the provided float64 value
func Float64(v float64) *float64 { return &v }

// Rune creates a pointer to the provided rune value
func Rune(v rune) *rune { return &v }

// Byte creates a pointer to the provided byte value
func Byte(v byte) *byte { return &v }

// String creates a pointer to the provided string value
func String(v string) *string { return &v }

// Complex64 creates a pointer to the provided complex64 value
func Complex64(v complex64) *complex64 { return &v }

// Complex128 creates a pointer to the provided complex128 value
func Complex128(v complex128) *complex128 { return &v }

// Interface creates a pointer to the provided interface{} value
func Interface(v interface{}) *interface{} { return &v }

// Duration creates a pointer to the provided time.Duration value
func Duration(v time.Duration) *time.Duration { return &v }

// Time creates a pointer to the provided time.Time value
func Time(v time.Time) *time.Time { return &v }
