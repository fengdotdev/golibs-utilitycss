package gorenderercss

import (
	"strings"

	"github.com/fengdotdev/golibs-utilitycss/v0/containers/gocss"
	"github.com/fengdotdev/golibs-utilitycss/v0/types"
)

var _ types.RendererCSS = (*GoRendererCSS)(nil)

type GoRendererCSS struct {

	// key is the class name
	// value is the css content
	// e.g. "bg-white" -> ".bg-white {background-color: #ffffff;}"
	content map[string]string
}

// Mix implements containers.RendererCSS.
func (g *GoRendererCSS) Mix(utilities ...ClassName) string {
	panic("unimplemented")
}

// Class implements containers.RendererCSS.
func (g *GoRendererCSS) Class(utilities ...string) string {
	panic("unimplemented")
}

// Render implements containers.RendererCSS.
func (g *GoRendererCSS) Render() CSS {
	return gocss.NewCss(g.render())
}

// Utility implements containers.RendererCSS.
func (g *GoRendererCSS) Utils(utilities ...Fn) string {

	acumulate := ""

	if len(utilities) == 0 {
		return ""
	}

	for _, fn := range utilities {

		utility := fn()
		class := utility.Class()
		css := utility.Css()
		g.content[class.String()] = css.String()

		acumulate += " " + class.String()
	}

	// remove the first space
	trim := strings.TrimSpace(acumulate)
	return trim
}

// String implements containers.CSS.
func (g *GoRendererCSS) String() string {
	return g.render()
}

func (g *GoRendererCSS) render() string {

	acumulate := ""

	for _, v := range g.content {
		acumulate += v + "\n"
	}
	return acumulate
}

func NewRendererCSS() *GoRendererCSS {
	return &GoRendererCSS{
		content: make(map[string]string),
	}
}
