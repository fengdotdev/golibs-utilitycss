package golibsutilitycss

type Selector string

func NewSelector(selector string) Selector {
	return Selector(selector)
}

func (s Selector) GetSelector() string {
	return string(s)
}
