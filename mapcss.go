package golibsutilitycss

import (
	"errors"
	"strings"
)

var (
	ErrSelectorAlreadyExist = errors.New("selector already exist")
	ErrSelectorNotFound     = errors.New("selector not found")
)

type MapCSSInterface interface {
	Add(selector Selector, declarations ...Declaration) error
	AddOverride(selector Selector, declarations ...Declaration)
	Remove(selector Selector) error
	Get(selector Selector) ([]Declaration, error)
	GetContent() map[Selector][]Declaration
	GetCSS() string
}

var _ MapCSSInterface = &MapCSS{}

type MapCSS struct {
	content map[Selector][]Declaration
}

func NewMapCSS() *MapCSS {
	return &MapCSS{
		content: make(map[Selector][]Declaration),
	}
}

func (m *MapCSS) Add(selector Selector, declarations ...Declaration) error {
	if _, ok := m.content[selector]; ok {
		return ErrSelectorAlreadyExist
	}
	m.content[selector] = declarations
	return nil
}

func (m *MapCSS) AddOverride(selector Selector, declarations ...Declaration) {
	m.content[selector] = declarations
}

func (m *MapCSS) Remove(selector Selector) error {
	if _, ok := m.content[selector]; !ok {
		return ErrSelectorNotFound
	}
	delete(m.content, selector)
	return nil
}

func (m *MapCSS) Get(selector Selector) ([]Declaration, error) {
	if _, ok := m.content[selector]; !ok {
		return nil, ErrSelectorNotFound
	}
	return m.content[selector], nil
}

func (m *MapCSS) GetContent() map[Selector][]Declaration {
	return m.content
}

func (m *MapCSS) GetCSS() string {
	var sb strings.Builder

	for selector, declarations := range m.content {
		sb.WriteString(selector.GetSelector())
		sb.WriteString(ConstOpenBracket)
		for _, declaration := range declarations {
			sb.WriteString(declaration.GetProperty())
			sb.WriteString(ConstSeparatorDeclararion)
			sb.WriteString(declaration.GetValue())
			sb.WriteString(ConstEndDeclaration)
		}
		sb.WriteString(ConstCloseBracket)
	}
	return sb.String()
}
