package helpers

import (
	"errors"

	"github.com/fengdotdev/golibs-utilitycss/v0/types"
)

// ForEach iterates over a Mapeable type (Set or List) and applies the function f to each element.
// It returns an error if the type is not supported.
func ForEach[T types.Mapeable](Mapeable T, f func(string)) error {
	switch v := any(Mapeable).(type) {
	case types.Set:
		ForEachSet(v, f)
		return nil
	case types.List:
		ForEachList(v, f)
		return nil
	default:
		return errors.New("unsupported type")
	}
}

// ForEachSet iterates over a Set and applies the function f to each key.
// It is a helper function to be used with ForEach.
// It does not return an error because it is expected to be used with a Set.
func ForEachSet(set types.Set, f func(string)) {
	for k := range set {
		f(k)
	}
}

// ForEachList iterates over a List and applies the function f to each element.
// It is a helper function to be used with ForEach.
// It does not return an error because it is expected to be used with a List.
func ForEachList(list types.List, f func(string)) {
	for _, k := range list {
		f(k)
	}
}
