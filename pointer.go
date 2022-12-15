package pointer

// New returns a new pointer to provided value.
func New[T any](val T) *T {
	return &val
}

// Value returns pointer value and true if not nil or zero value and false.
func Value[T any](ptr *T) (T, bool) {
	if ptr != nil {
		return *ptr, true
	}

	var zeroValue T

	return zeroValue, false
}

// Coalesce returns the first non-nil value in a list (or fallback value when all nils).
func Coalesce[T any](fallback T, values ...*T) T {
	for _, value := range values {
		if value != nil {
			return *value
		}
	}

	return fallback
}
