package types


// Set is a type alias for a map with string keys and empty struct values.
// It is used to represent a set of unique strings, where the value is not important.
type Set = map[string]struct{}

