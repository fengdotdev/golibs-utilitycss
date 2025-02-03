package golibsutilitycss

import (
	"os"
)

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

func ReadFileRemote(url string) (string, error) {
	return "", nil
}

func ReadFile(cssfile string) (string, error) {

	file, err := os.Open(cssfile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func GetCode(css string) (map[string]PairKV, error) {
	// regex to get the css code

	cursor := 0
	totallen := len(css)

	for cursor < totallen {
		// get the selector
		// get the property
		// get the value
		// add to the map
	}

	return nil, nil
}
