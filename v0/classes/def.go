package classes

import (
	"github.com/fengdotdev/golibs-utilitycss/v0/containers/goclass"
	"github.com/fengdotdev/golibs-utilitycss/v0/helpers"
	"github.com/fengdotdev/golibs-utilitycss/v0/types"
)

type Utility = types.Utility

type GoClass = goclass.GoClass

func NewUtility(class string, css string) Utility {
	return helpers.NewUtility(class, css)
}
