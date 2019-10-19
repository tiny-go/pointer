package pointer

import "testing"

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
				return BoolVal(ptr.(*bool))
			},
			input:  func() *bool { return nil }(),
			output: false,
			valid:  false,
		},
		{
			title: "test BoolVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return BoolVal(ptr.(*bool))
			},
			input:  func() *bool { b := true; return &b }(),
			output: true,
			valid:  true,
		},
		{
			title: "test UintVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return UintVal(ptr.(*uint))
			},
			input:  func() *uint { return nil }(),
			output: uint(0),
			valid:  false,
		},
		{
			title: "test UintVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return UintVal(ptr.(*uint))
			},
			input:  func() *uint { b := uint(42); return &b }(),
			output: uint(42),
			valid:  true,
		},
		{
			title: "test Uint8Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Uint8Val(ptr.(*uint8))
			},
			input:  func() *uint8 { return nil }(),
			output: uint8(0),
			valid:  false,
		},
		{
			title: "test Uint8Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Uint8Val(ptr.(*uint8))
			},
			input:  func() *uint8 { b := uint8(42); return &b }(),
			output: uint8(42),
			valid:  true,
		},
		{
			title: "test Uint16Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Uint16Val(ptr.(*uint16))
			},
			input:  func() *uint16 { return nil }(),
			output: uint16(0),
			valid:  false,
		},
		{
			title: "test Uint16Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Uint16Val(ptr.(*uint16))
			},
			input:  func() *uint16 { b := uint16(42); return &b }(),
			output: uint16(42),
			valid:  true,
		},
		{
			title: "test Uint32Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Uint32Val(ptr.(*uint32))
			},
			input:  func() *uint32 { return nil }(),
			output: uint32(0),
			valid:  false,
		},
		{
			title: "test Uint32Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Uint32Val(ptr.(*uint32))
			},
			input:  func() *uint32 { b := uint32(42); return &b }(),
			output: uint32(42),
			valid:  true,
		},
		{
			title: "test Uint64Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Uint64Val(ptr.(*uint64))
			},
			input:  func() *uint64 { return nil }(),
			output: uint64(0),
			valid:  false,
		},
		{
			title: "test Uint64Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Uint64Val(ptr.(*uint64))
			},
			input:  func() *uint64 { b := uint64(42); return &b }(),
			output: uint64(42),
			valid:  true,
		},
		{
			title: "test UintptrVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return UintptrVal(ptr.(*uintptr))
			},
			input:  func() *uintptr { return nil }(),
			output: uintptr(0),
			valid:  false,
		},
		{
			title: "test UintptrVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return UintptrVal(ptr.(*uintptr))
			},
			input:  func() *uintptr { b := uintptr(42); return &b }(),
			output: uintptr(42),
			valid:  true,
		},
		{
			title: "test IntVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return IntVal(ptr.(*int))
			},
			input:  func() *int { return nil }(),
			output: 0,
			valid:  false,
		},
		{
			title: "test IntVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return IntVal(ptr.(*int))
			},
			input:  func() *int { b := 42; return &b }(),
			output: 42,
			valid:  true,
		},
		{
			title: "test Int8Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Int8Val(ptr.(*int8))
			},
			input:  func() *int8 { return nil }(),
			output: int8(0),
			valid:  false,
		},
		{
			title: "test Int8Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Int8Val(ptr.(*int8))
			},
			input:  func() *int8 { b := int8(42); return &b }(),
			output: int8(42),
			valid:  true,
		},
		{
			title: "test Int16Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Int16Val(ptr.(*int16))
			},
			input:  func() *int16 { return nil }(),
			output: int16(0),
			valid:  false,
		},
		{
			title: "test Int16Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Int16Val(ptr.(*int16))
			},
			input:  func() *int16 { b := int16(42); return &b }(),
			output: int16(42),
			valid:  true,
		},
		{
			title: "test Int32Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Int32Val(ptr.(*int32))
			},
			input:  func() *int32 { return nil }(),
			output: int32(0),
			valid:  false,
		},
		{
			title: "test Int32Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Int32Val(ptr.(*int32))
			},
			input:  func() *int32 { b := int32(42); return &b }(),
			output: int32(42),
			valid:  true,
		},
		{
			title: "test Int64Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Int64Val(ptr.(*int64))
			},
			input:  func() *int64 { return nil }(),
			output: int64(0),
			valid:  false,
		},
		{
			title: "test Int64Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Int64Val(ptr.(*int64))
			},
			input:  func() *int64 { b := int64(42); return &b }(),
			output: int64(42),
			valid:  true,
		},
		{
			title: "test Float32Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Float32Val(ptr.(*float32))
			},
			input:  func() *float32 { return nil }(),
			output: float32(0),
			valid:  false,
		},
		{
			title: "test Float32Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Float32Val(ptr.(*float32))
			},
			input:  func() *float32 { b := float32(42); return &b }(),
			output: float32(42),
			valid:  true,
		},
		{
			title: "test Float64Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Float64Val(ptr.(*float64))
			},
			input:  func() *float64 { return nil }(),
			output: float64(0),
			valid:  false,
		},
		{
			title: "test Float64Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Float64Val(ptr.(*float64))
			},
			input:  func() *float64 { b := float64(42); return &b }(),
			output: float64(42),
			valid:  true,
		},
		{
			title: "test RuneVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return RuneVal(ptr.(*rune))
			},
			input:  func() *rune { return nil }(),
			output: rune(0),
			valid:  false,
		},
		{
			title: "test RuneVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return RuneVal(ptr.(*rune))
			},
			input:  func() *rune { b := rune(42); return &b }(),
			output: rune(42),
			valid:  true,
		},
		{
			title: "test ByteVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return ByteVal(ptr.(*byte))
			},
			input:  func() *byte { return nil }(),
			output: byte(0),
			valid:  false,
		},
		{
			title: "test ByteVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return ByteVal(ptr.(*byte))
			},
			input:  func() *byte { b := byte(42); return &b }(),
			output: byte(42),
			valid:  true,
		},
		{
			title: "test StringVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return StringVal(ptr.(*string))
			},
			input:  func() *string { return nil }(),
			output: "",
			valid:  false,
		},
		{
			title: "test StringVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return StringVal(ptr.(*string))
			},
			input:  func() *string { b := "42"; return &b }(),
			output: "42",
			valid:  true,
		},
		{
			title: "test Complex64Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Complex64Val(ptr.(*complex64))
			},
			input:  func() *complex64 { return nil }(),
			output: complex64(0),
			valid:  false,
		},
		{
			title: "test Complex64Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Complex64Val(ptr.(*complex64))
			},
			input:  func() *complex64 { b := complex64(42); return &b }(),
			output: complex64(42),
			valid:  true,
		},
		{
			title: "test Complex128Val providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Complex128Val(ptr.(*complex128))
			},
			input:  func() *complex128 { return nil }(),
			output: complex128(0),
			valid:  false,
		},
		{
			title: "test Complex128Val providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return Complex128Val(ptr.(*complex128))
			},
			input:  func() *complex128 { b := complex128(42); return &b }(),
			output: complex128(42),
			valid:  true,
		},
		{
			title: "test InterfaceVal providing `nil` pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return InterfaceVal(ptr.(*interface{}))
			},
			input:  func() *interface{} { return nil }(),
			output: nil,
			valid:  false,
		},
		{
			title: "test InterfaceVal providing valid pointer",
			fn: func(ptr interface{}) (interface{}, bool) {
				return InterfaceVal(ptr.(*interface{}))
			},
			input:  func() *interface{} { b := interface{}(42); return &b }(),
			output: 42,
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
