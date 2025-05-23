package classes

import (
	"fmt"

	"github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/containers"
	"github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/containers/goclass"
	"github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/containers/gocss"
	"github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/containers/goutility"
)

type Utility = containers.Utility

func NewUtility(class string, css string) Utility {

	classObj := goclass.NewClass(class)
	cssObj := gocss.NewCss(css)

	return goutility.NewUtility(cssObj, classObj)
}

func NewSelectorBlock(selector string, css string) string {

	return fmt.Sprintf(".%s { %s }", selector, css)
}

func Bg_white() Utility {
	css := "background-color: #ffffff;"
	class := "bg-white"
	block := NewSelectorBlock(class, css)
	return NewUtility(class, block)
}

func Bg_black() Utility {
	css := "background-color: #000000;"
	class := "bg-black"
	block := NewSelectorBlock(class, css)
	return NewUtility(class, block)
}
