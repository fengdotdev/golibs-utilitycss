package goutility

import "github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/containers"

var _ containers.Utility = (*GoUtility)(nil)

func NewUtility(css containers.CSS, class containers.Class) *GoUtility {
	return &GoUtility{
		css:   css,
		class: class,
	}
}

type GoUtility struct {
	css   containers.CSS
	class containers.Class
}

// Class implements containers.Utility.
func (g *GoUtility) Class() containers.Class {
	return g.class
}

// Css implements containers.Utility.
func (g *GoUtility) Css() containers.CSS {
	return g.css
}
