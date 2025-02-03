package golibsutilitycss

func NewDeclarations() []Declaration {
	return make([]Declaration, 0)
}

type Declaration struct {
	property string
	value    string
}

func NewPair(property string, value string) Declaration {
	return NewDeclaration(property, value)
}

func NewDeclaration(property string, value string) Declaration {
	return Declaration{
		property: property,
		value:    value,
	}
}

func (d *Declaration) GetProperty() string {
	return d.property
}

func (d *Declaration) GetValue() string {
	return d.value
}
