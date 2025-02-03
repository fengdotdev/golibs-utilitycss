package golibsutilitycss

import "strings"

type MapCSS struct {
	content map[Selector]Declaration
}

func NewMapCSS() *MapCSS {
	return &MapCSS{
		content: make(map[Selector]Declaration),
	}
}

func (m *MapCSS) Add(selector Selector, declaration Declaration) {
	m.content[selector] = declaration
}

func (m *MapCSS) Get(selector Selector) Declaration {
	return m.content[selector]
}

func (m *MapCSS) GetContent() map[Selector]Declaration {
	return m.content
}

func (m *MapCSS) GetCSS() (string, error) {

	var css strings.Builder
	var err error

	for selector, declaration := range m.content {
		csscode := CSSCode(selector.GetSelector(), declaration.GetProperty(), declaration.GetValue())
		_, err = css.WriteString(csscode)
		if err != nil {
			break
		}
	}
	if err != nil {
		return "", err
	}
	return css.String(), nil
}

