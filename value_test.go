package pointer_test

import (
	"testing"
	"time"

	"github.com/tiny-go/pointer"
)

func Test_Value(t *testing.T) {
	type testCase struct {
		title  string
		fn     func(interface{}) (interface{}, bool)
		input  interface{}
		output interface{}
		valid  bool
	}

	testCases := []testCase{
		{
			title: "test BoolVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.BoolVal(ptr.(*bool))
			},
			input:  func() *bool { return nil }(),
			output: false,
			valid:  false,
		},
		{
			title: "test BoolVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.BoolVal(ptr.(*bool))
			},
			input:  func() *bool { b := true; return &b }(),
			output: true,
			valid:  true,
		},
		{
			title: "test UintVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.UintVal(ptr.(*uint))
			},
			input:  func() *uint { return nil }(),
			output: uint(0),
			valid:  false,
		},
		{
			title: "test UintVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.UintVal(ptr.(*uint))
			},
			input:  func() *uint { u := uint(42); return &u }(),
			output: uint(42),
			valid:  true,
		},
		{
			title: "test Uint8Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Uint8Val(ptr.(*uint8))
			},
			input:  func() *uint8 { return nil }(),
			output: uint8(0),
			valid:  false,
		},
		{
			title: "test Uint8Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Uint8Val(ptr.(*uint8))
			},
			input:  func() *uint8 { u := uint8(42); return &u }(),
			output: uint8(42),
			valid:  true,
		},
		{
			title: "test Uint16Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Uint16Val(ptr.(*uint16))
			},
			input:  func() *uint16 { return nil }(),
			output: uint16(0),
			valid:  false,
		},
		{
			title: "test Uint16Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Uint16Val(ptr.(*uint16))
			},
			input:  func() *uint16 { u := uint16(42); return &u }(),
			output: uint16(42),
			valid:  true,
		},
		{
			title: "test Uint32Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Uint32Val(ptr.(*uint32))
			},
			input:  func() *uint32 { return nil }(),
			output: uint32(0),
			valid:  false,
		},
		{
			title: "test Uint32Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Uint32Val(ptr.(*uint32))
			},
			input:  func() *uint32 { u := uint32(42); return &u }(),
			output: uint32(42),
			valid:  true,
		},
		{
			title: "test Uint64Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Uint64Val(ptr.(*uint64))
			},
			input:  func() *uint64 { return nil }(),
			output: uint64(0),
			valid:  false,
		},
		{
			title: "test Uint64Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Uint64Val(ptr.(*uint64))
			},
			input:  func() *uint64 { u := uint64(42); return &u }(),
			output: uint64(42),
			valid:  true,
		},
		{
			title: "test UintptrVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.UintptrVal(ptr.(*uintptr))
			},
			input:  func() *uintptr { return nil }(),
			output: uintptr(0),
			valid:  false,
		},
		{
			title: "test UintptrVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.UintptrVal(ptr.(*uintptr))
			},
			input:  func() *uintptr { p := uintptr(42); return &p }(),
			output: uintptr(42),
			valid:  true,
		},
		{
			title: "test IntVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.IntVal(ptr.(*int))
			},
			input:  func() *int { return nil }(),
			output: 0,
			valid:  false,
		},
		{
			title: "test IntVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.IntVal(ptr.(*int))
			},
			input:  func() *int { i := 42; return &i }(),
			output: 42,
			valid:  true,
		},
		{
			title: "test Int8Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Int8Val(ptr.(*int8))
			},
			input:  func() *int8 { return nil }(),
			output: int8(0),
			valid:  false,
		},
		{
			title: "test Int8Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Int8Val(ptr.(*int8))
			},
			input:  func() *int8 { i := int8(42); return &i }(),
			output: int8(42),
			valid:  true,
		},
		{
			title: "test Int16Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Int16Val(ptr.(*int16))
			},
			input:  func() *int16 { return nil }(),
			output: int16(0),
			valid:  false,
		},
		{
			title: "test Int16Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Int16Val(ptr.(*int16))
			},
			input:  func() *int16 { i := int16(42); return &i }(),
			output: int16(42),
			valid:  true,
		},
		{
			title: "test Int32Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Int32Val(ptr.(*int32))
			},
			input:  func() *int32 { return nil }(),
			output: int32(0),
			valid:  false,
		},
		{
			title: "test Int32Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Int32Val(ptr.(*int32))
			},
			input:  func() *int32 { i := int32(42); return &i }(),
			output: int32(42),
			valid:  true,
		},
		{
			title: "test Int64Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Int64Val(ptr.(*int64))
			},
			input:  func() *int64 { return nil }(),
			output: int64(0),
			valid:  false,
		},
		{
			title: "test Int64Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Int64Val(ptr.(*int64))
			},
			input:  func() *int64 { i := int64(42); return &i }(),
			output: int64(42),
			valid:  true,
		},
		{
			title: "test Float32Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Float32Val(ptr.(*float32))
			},
			input:  func() *float32 { return nil }(),
			output: float32(0),
			valid:  false,
		},
		{
			title: "test Float32Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Float32Val(ptr.(*float32))
			},
			input:  func() *float32 { f := float32(42); return &f }(),
			output: float32(42),
			valid:  true,
		},
		{
			title: "test Float64Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Float64Val(ptr.(*float64))
			},
			input:  func() *float64 { return nil }(),
			output: float64(0),
			valid:  false,
		},
		{
			title: "test Float64Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Float64Val(ptr.(*float64))
			},
			input:  func() *float64 { f := float64(42); return &f }(),
			output: float64(42),
			valid:  true,
		},
		{
			title: "test RuneVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.RuneVal(ptr.(*rune))
			},
			input:  func() *rune { return nil }(),
			output: rune(0),
			valid:  false,
		},
		{
			title: "test RuneVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.RuneVal(ptr.(*rune))
			},
			input:  func() *rune { r := rune(42); return &r }(),
			output: rune(42),
			valid:  true,
		},
		{
			title: "test ByteVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.ByteVal(ptr.(*byte))
			},
			input:  func() *byte { return nil }(),
			output: byte(0),
			valid:  false,
		},
		{
			title: "test ByteVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.ByteVal(ptr.(*byte))
			},
			input:  func() *byte { b := byte(42); return &b }(),
			output: byte(42),
			valid:  true,
		},
		{
			title: "test StringVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.StringVal(ptr.(*string))
			},
			input:  func() *string { return nil }(),
			output: "",
			valid:  false,
		},
		{
			title: "test StringVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.StringVal(ptr.(*string))
			},
			input:  func() *string { s := "42"; return &s }(),
			output: "42",
			valid:  true,
		},
		{
			title: "test Complex64Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Complex64Val(ptr.(*complex64))
			},
			input:  func() *complex64 { return nil }(),
			output: complex64(0),
			valid:  false,
		},
		{
			title: "test Complex64Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Complex64Val(ptr.(*complex64))
			},
			input:  func() *complex64 { c := complex64(42); return &c }(),
			output: complex64(42),
			valid:  true,
		},
		{
			title: "test Complex128Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Complex128Val(ptr.(*complex128))
			},
			input:  func() *complex128 { return nil }(),
			output: complex128(0),
			valid:  false,
		},
		{
			title: "test Complex128Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.Complex128Val(ptr.(*complex128))
			},
			input:  func() *complex128 { c := complex128(42); return &c }(),
			output: complex128(42),
			valid:  true,
		},
		{
			title: "test InterfaceVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.InterfaceVal(ptr.(*interface{}))
			},
			input:  func() *interface{} { return nil }(),
			output: nil,
			valid:  false,
		},
		{
			title: "test InterfaceVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.InterfaceVal(ptr.(*interface{}))
			},
			input:  func() *interface{} { i := interface{}(42); return &i }(),
			output: 42,
			valid:  true,
		},
		{
			title: "test DurationVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.DurationVal(ptr.(*time.Duration))
			},
			input:  func() *time.Duration { return nil }(),
			output: time.Duration(0),
			valid:  false,
		},
		{
			title: "test DurationVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.DurationVal(ptr.(*time.Duration))
			},
			input:  func() *time.Duration { d := time.Duration(42); return &d }(),
			output: time.Duration(42),
			valid:  true,
		},
		{
			title: "test TimeVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.TimeVal(ptr.(*time.Time))
			},
			input:  func() *time.Time { return nil }(),
			output: time.Time{},
			valid:  false,
		},
		{
			title: "test TimeVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return pointer.TimeVal(ptr.(*time.Time))
			},
			input:  func() *time.Time { t := time.Time{}; return &t }(),
			output: time.Time{},
			valid:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			val, ok := tc.fn(tc.input)
			if val != tc.output {
				t.Errorf("value should equal `%v`", tc.output)
			}
			if ok != tc.valid {
				t.Errorf("`ok` was expected to be `%v`", tc.valid)
			}
		})
	}
}
