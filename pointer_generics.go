package pointer

// Pointer returns a pointer to provided value.
func Pointer[T any](val T) *T {
	return &val
}
