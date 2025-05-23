package containers

// constraint ClassName to be either string or Utility
type ClassName interface{}

type RendererCSS interface {
	Render() CSS

	Mix(utilities ...ClassName) string
	Class(utilities ...string) string
	Utils(utilities ...Fn) string
}
