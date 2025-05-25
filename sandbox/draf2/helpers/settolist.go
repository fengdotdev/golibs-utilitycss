package helpers

import "github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/types"

func SetToList(set types.Set) types.List {
	list := make([]string, 0, len(set))
	for k := range set {
		list = append(list, k)
	}
	return list
}

func ListToSet(list types.List) types.Set {
	set := make(types.Set)
	for _, k := range list {
		set[k] = struct{}{}
	}
	return set
}
