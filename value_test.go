package pointer

import "testing"

func Test_BoolVal(t *testing.T) {
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
