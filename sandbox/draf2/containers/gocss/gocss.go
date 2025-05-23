package gocss

import "github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/containers"

var _ containers.CSS = (*GoCSS)(nil)

func NewCss(content string) *GoCSS {
	return &GoCSS{
		content: content,
	}
}

type GoCSS struct {
	content string
}

// String implements containers.Css.
func (g *GoCSS) String() string {
	return g.content
}
