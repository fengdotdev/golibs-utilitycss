package helpers

import (
	"errors"

	"github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/types"
)

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

func ForEachSet(set types.Set, f func(string)) {
	for k := range set {
		f(k)
	}
}

func ForEachList(list types.List, f func(string)) {
	for _, k := range list {
		f(k)
	}
}
