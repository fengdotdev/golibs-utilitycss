package golibsutilitycss

import "strings"

type Selector string

func NewSelector(selector string) Selector {
	return Selector(selector)
}

func (s Selector) GetSelector() string {
	return string(s)
}

func (s Selector) String() string {
	return string(s)
}

func (s Selector) Equals(selector Selector) bool {
	return s.GetSelector() == selector.GetSelector()
}

func (s Selector) IsEmpty() bool {
	return s.GetSelector() == ""
}

func (s Selector) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s Selector) IsEqual(selector Selector) bool {
	return s.GetSelector() == selector.GetSelector()
}

func (s Selector) NumberOfSelector() int {
	return len(strings.Split(s.GetSelector(), ConstSeparatorSelector))
}
