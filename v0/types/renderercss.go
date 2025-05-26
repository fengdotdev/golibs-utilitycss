package types



// RendererCSS is an interface that defines methods for rendering CSS,
type RendererCSS interface {
	Render() CSS
	Mix(utilities ...ClassName) string  
	Class(utilities ...string) string
	Utils(utilities ...Ufunc) string
}
