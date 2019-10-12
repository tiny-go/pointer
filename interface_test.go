package pointer

import "testing"

func Test_Unpoint(t *testing.T) {
	t.Run("test if `nil` is returned providing nil", func(t *testing.T) {
		if Unpoint(nil) != nil {
			t.Error("unpointed value does not match expected value")
		}
	})

	t.Run("test if `nil` is returned providing nil pointer", func(t *testing.T) {
		var nilPointer *int
		if Unpoint(nilPointer) != nil {
			t.Error("unpointed value does not match expected value")
		}
	})

	t.Run("test if pointer value is returned", func(t *testing.T) {
		intVal := 42
		ptr := &intVal
		if Unpoint(ptr) != intVal {
			t.Error("unpointed value does not match expected value")
		}
	})
}
