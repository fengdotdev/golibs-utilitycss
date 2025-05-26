package gocss

var _ CSS = (*GoCSS)(nil)

func NewCss(content string) CSS {
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
