package helpers

import "github.com/fengdotdev/golibs-utilitycss/v0/types"


// SetToList converts a Set to a List.
// It iterates over the keys of the Set and appends them to a new List.
// The resulting List contains all unique keys from the Set.
func SetToList(set types.Set) types.List {
	list := make([]string, 0, len(set))
	for k := range set {
		list = append(list, k)
	}
	return list
}

// ListToSet converts a List to a Set.
// It iterates over the elements of the List and adds them to a new Set.
// The resulting Set contains all unique elements from the List.
// If the List contains duplicate elements, they will be ignored in the Set.
// This is useful for ensuring that the Set contains only unique elements.
func ListToSet(list types.List) types.Set {
	set := make(types.Set)
	for _, k := range list {
		set[k] = struct{}{}
	}
	return set
}
