package golibsutilitycss

func CSSCode(selector, property, value string) string {
	return selector + " { " + property + ": " + value + "; }"
}

func Populate(m *MapCSS) error {

	m.Add(NewSelector("h1"), NewDeclaration("font-size", "2em"))

	return nil
}

func PopulateFromMap(m *MapCSS, cssfile string) error {

	return nil
}

func ReadFile(cssfile string) (string, error) {
	return "", nil
}

func GetCode(css string) (map[string]PairKV, error) {
	return nil, nil
}
