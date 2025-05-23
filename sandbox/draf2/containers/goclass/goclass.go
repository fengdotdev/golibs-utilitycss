package goclass

import "github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/containers"

var _ containers.Class = (*GoClass)(nil)

func NewClass(content string) *GoClass {
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
