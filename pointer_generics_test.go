package pointer_test

import (
	"testing"

	"github.com/tiny-go/pointer"
)

func Test_Pointer_Generics(t *testing.T) {
	t.Run("Pointer", func(t *testing.T) {
		value := "foo"

		if *pointer.Pointer(value) != value {
			t.Errorf("unexpected value: %v", value)
		}
	})
}
