package types


// Mapeable is a type constraint that allows a type to be either 
// a map with string keys and struct values or a slice of strings.
// It is used to define types that can be mapped or iterated over in a consistent way.
// It is useful for functions that need to accept either a map or a slice of strings as input.
type Mapeable interface {
	map[string]struct{} | []string
}
