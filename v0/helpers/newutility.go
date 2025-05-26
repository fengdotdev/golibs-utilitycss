package helpers

import (
	"github.com/fengdotdev/golibs-utilitycss/v0/containers/goclass"
	"github.com/fengdotdev/golibs-utilitycss/v0/containers/gocss"
	"github.com/fengdotdev/golibs-utilitycss/v0/containers/goutility"
	"github.com/fengdotdev/golibs-utilitycss/v0/types"
)

// NewUtility creates a new Utility instance with the given class and CSS.
// It initializes a Class object and a Css object, then returns a Utility object.
// Example: NewUtility("bg-white", "background-color: #ffffff;")
func NewUtility(class string, css string) types.Utility {

	classObj := goclass.NewClass(class)
	block := NewSelectorBlock(class, css)
	cssObj := gocss.NewCss(block)

	return goutility.NewUtility(cssObj, classObj)
}
