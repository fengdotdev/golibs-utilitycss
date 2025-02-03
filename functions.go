package golibsutilitycss

import (
	"os"
	"strings"
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


// example bg-white{background-color: var(--color-white);}  -> bg-white, background-color, white
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

// example  flex flex-col items-center  -> [flex, flex-col, items-center]
func StringToSelectors(utilityClasses string) ([]Selector, error) {

	utilityClasses = strings.TrimSpace(utilityClasses)
	if utilityClasses == "" {
		// Return an empty slice if the input is empty.
		return []Selector{}, nil
	}

	// Use strings.Fields to split by any whitespace.
	parts := strings.Fields(utilityClasses)

	// Create a slice of selectors with the same length as the parts.
	selectors := make([]Selector, len(parts))
	for i, part := range parts {
		selectors[i] = Selector(part)
	}

	return selectors, nil
}
