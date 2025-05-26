package goutility

var _ Utility = (*GoUtility)(nil)

func NewUtility(css CSS, class Class) Utility {
	return &GoUtility{
		css:   css,
		class: class,
	}
}

type GoUtility struct {
	css   CSS
	class Class
}

// Class implements containers.Utility.
func (g *GoUtility) Class() Class {
	return g.class
}

// Css implements containers.Utility.
func (g *GoUtility) Css() CSS {
	return g.css
}
