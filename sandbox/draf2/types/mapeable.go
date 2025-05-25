package types

type Mapeable interface {
	map[string]struct{} | []string
}
