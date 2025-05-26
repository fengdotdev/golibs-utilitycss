package helpers

// GetFirstElement returns the first element of a slice and a boolean indicating if the slice is empty.
// If the slice is empty, it returns the zero value of the element type and false.
func GetFirstElement[T any](s []T) (T, bool) {
	if len(s) > 0 {
		return s[0], true
	}

	var zeroValue T
	return zeroValue, false
}

// GetFirstString returns the first string of a slice and a boolean indicating if the slice is empty.
// If the slice is empty, it returns an empty string and false.
func GetFirstString(s []string) (string, bool) {
	return GetFirstElement(s)
}
