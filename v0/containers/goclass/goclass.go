package goclass

var _ Class = (*GoClass)(nil)

func NewClass(content string) CSS {
	return &GoClass{
		content: content,
	}
}

type GoClass struct {
	content string
}

// String implements containers.Class.
func (g *GoClass) String() string {
	return g.content
}
