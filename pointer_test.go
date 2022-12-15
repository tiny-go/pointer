package pointer_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/tiny-go/pointer"
)

func Test_Pointer(t *testing.T) {
	type testCase struct {
		value  interface{}
		fn     func(interface{}) interface{}
		output interface{}
	}

	testCases := []testCase{
		{
			value:  true,
			fn:     func(v interface{}) interface{} { return pointer.Bool(v.(bool)) },
			output: func() interface{} { b := true; return &b }(),
		},
		{
			value:  uint(42),
			fn:     func(v interface{}) interface{} { return pointer.Uint(v.(uint)) },
			output: func() interface{} { i := uint(42); return &i }(),
		},
		{
			value:  uint8(42),
			fn:     func(v interface{}) interface{} { return pointer.Uint8(v.(uint8)) },
			output: func() interface{} { i := uint8(42); return &i }(),
		},
		{
			value:  uint16(42),
			fn:     func(v interface{}) interface{} { return pointer.Uint16(v.(uint16)) },
			output: func() interface{} { i := uint16(42); return &i }(),
		},
		{
			value:  uint32(42),
			fn:     func(v interface{}) interface{} { return pointer.Uint32(v.(uint32)) },
			output: func() interface{} { i := uint32(42); return &i }(),
		},
		{
			value:  uint64(42),
			fn:     func(v interface{}) interface{} { return pointer.Uint64(v.(uint64)) },
			output: func() interface{} { i := uint64(42); return &i }(),
		},
		{
			value:  42,
			fn:     func(v interface{}) interface{} { return pointer.Int(v.(int)) },
			output: func() interface{} { i := 42; return &i }(),
		},
		{
			value:  int8(42),
			fn:     func(v interface{}) interface{} { return pointer.Int8(v.(int8)) },
			output: func() interface{} { i := int8(42); return &i }(),
		},
		{
			value:  int16(42),
			fn:     func(v interface{}) interface{} { return pointer.Int16(v.(int16)) },
			output: func() interface{} { i := int16(42); return &i }(),
		},
		{
			value:  int32(42),
			fn:     func(v interface{}) interface{} { return pointer.Int32(v.(int32)) },
			output: func() interface{} { i := int32(42); return &i }(),
		},
		{
			value:  int64(42),
			fn:     func(v interface{}) interface{} { return pointer.Int64(v.(int64)) },
			output: func() interface{} { i := int64(42); return &i }(),
		},
		{
			value:  uintptr(42),
			fn:     func(v interface{}) interface{} { return pointer.Uintptr(v.(uintptr)) },
			output: func() interface{} { i := uintptr(42); return &i }(),
		},
		{
			value:  float32(42),
			fn:     func(v interface{}) interface{} { return pointer.Float32(v.(float32)) },
			output: func() interface{} { i := float32(42); return &i }(),
		},
		{
			value:  float64(42),
			fn:     func(v interface{}) interface{} { return pointer.Float64(v.(float64)) },
			output: func() interface{} { i := float64(42); return &i }(),
		},
		{
			value:  rune(42),
			fn:     func(v interface{}) interface{} { return pointer.Rune(v.(rune)) },
			output: func() interface{} { i := rune(42); return &i }(),
		},
		{
			value:  byte(42),
			fn:     func(v interface{}) interface{} { return pointer.Byte(v.(byte)) },
			output: func() interface{} { i := byte(42); return &i }(),
		},
		{
			value:  "42",
			fn:     func(v interface{}) interface{} { return pointer.String(v.(string)) },
			output: func() interface{} { i := "42"; return &i }(),
		},
		{
			value:  complex64(42),
			fn:     func(v interface{}) interface{} { return pointer.Complex64(v.(complex64)) },
			output: func() interface{} { i := complex64(42); return &i }(),
		},
		{
			value:  complex128(42),
			fn:     func(v interface{}) interface{} { return pointer.Complex128(v.(complex128)) },
			output: func() interface{} { i := complex128(42); return &i }(),
		},
		{
			value:  interface{}(42),
			fn:     func(v interface{}) interface{} { return pointer.Interface(v) },
			output: func() interface{} { i := interface{}(42); return &i }(),
		},
		{
			value:  time.Duration(42),
			fn:     func(v interface{}) interface{} { return pointer.Duration(v.(time.Duration)) },
			output: func() interface{} { i := time.Duration(42); return &i }(),
		},
		{
			value:  time.Time{},
			fn:     func(v interface{}) interface{} { return pointer.Time(v.(time.Time)) },
			output: func() interface{} { i := time.Time{}; return &i }(),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("test getting a pointer to `%T`", tc.value), func(t *testing.T) {
			if !reflect.DeepEqual(tc.fn(tc.value), tc.output) {
				t.Errorf("value should equal `%v`", tc.output)
			}
		})
	}
}
