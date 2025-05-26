package helpers

import (
	"fmt"
	"strings"

	"github.com/fengdotdev/golibs-utilitycss/v0/types"
	"golang.org/x/net/html"
)

// Parser parses the HTML content and extracts all class names.
// It returns a set of class names found
// uses a goroutine to find classes in the HTML document.
func Parser(htmlContent string) (types.Set, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {

		return nil, fmt.Errorf("error parsing HTML: %w", err)
	}

	// search for classes

	finds := make(chan string)

	go func() {
		FinderClasses(doc, finds)
		close(finds)

	}()

	classes := make(types.Set)
	for class := range finds {
		// Add each class to the set
		classes[class] = struct{}{}
	}
	return classes, nil
}


// ParserWithCallback parses the HTML content and calls the provided callback function for each class found.
// This function is useful for processing classes without using channels.
func ParserWithCallback(htmlContent string, callback func(class string)) error {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return fmt.Errorf("error parsing HTML: %w", err)
	}

	// search for classes
	FinderClassesWithCallback(doc, callback)

	return nil
}
