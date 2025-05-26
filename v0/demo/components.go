package demo

// Component is the interface that all components must implement.
type Component interface {

	// Children returns the child components of this component.
	Children() []Component

	// Render renders the component to a string.
	Render() (string, error)

	// Content returns the content of the component.
	Content() string

	// Class returns html class of the component.
	Class() string
}
