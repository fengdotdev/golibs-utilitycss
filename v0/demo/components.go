package demo

type Component interface {
	Children() []Component
	Render() (string, error)
	Content() string
	Class() string
}
