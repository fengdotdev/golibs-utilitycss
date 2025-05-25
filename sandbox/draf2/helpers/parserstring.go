package helpers

import (
	"fmt"
	"strings"

	"github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/types"
	"golang.org/x/net/html"
)

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
