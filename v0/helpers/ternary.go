package helpers

// Ternary returns trueValue if condition is true, otherwise returns falseValue.
// It is a generic function that works with any type T.
// This is useful for simplifying conditional logic in your code.
func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
