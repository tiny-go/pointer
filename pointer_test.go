package pointer_test

import (
	"testing"

	"github.com/tiny-go/pointer"
)

func Test_New(t *testing.T) {
	t.Run("Pointer", func(t *testing.T) {
		value := "foo"

		if *pointer.New(value) != value {
			t.Errorf("unexpected value: %v", value)
		}
	})
}

func Test_Value(t *testing.T) {
	t.Run("Scalar type (nil)", func(t *testing.T) {
		var ptr *float64

		value, ok := pointer.Value(ptr)
		if ok {
			t.Errorf("`ok` was expected to be false")
		}
		if value != 0 {
			t.Errorf("unexpected value: %v", value)
		}
	})

	t.Run("Scalar type (non nil)", func(t *testing.T) {
		var v float64 = 3.14

		value, ok := pointer.Value(&v)
		if !ok {
			t.Errorf("`ok` was expected to be true")
		}
		if value != 3.14 {
			t.Errorf("unexpected value: %v", value)
		}
	})

	t.Run("Interface (nil)", func(t *testing.T) {
		var ptr *interface{}

		value, ok := pointer.Value(ptr)
		if ok {
			t.Errorf("`ok` was expected to be false")
		}
		if value != nil {
			t.Errorf("value was expected to be nil")
		}
	})

	t.Run("Interface (non nil)", func(t *testing.T) {
		var v interface{} = "foo"

		value, ok := pointer.Value(&v)
		if !ok {
			t.Errorf("`ok` was expected to be true")
		}
		if value != "foo" {
			t.Errorf("unexpected value: %v", value)
		}
	})
}

func Test_Coalesce(t *testing.T) {
	t.Run("From collection", func(t *testing.T) {
		if value := pointer.Coalesce(1, nil, nil, pointer.New(42)); value != 42 {
			t.Errorf("unexpected value: %v", value)
		}
	})

	t.Run("Fallback", func(t *testing.T) {
		if value := pointer.Coalesce(1, nil, nil, nil); value != 1 {
			t.Errorf("unexpected value: %v", value)
		}
	})
}
