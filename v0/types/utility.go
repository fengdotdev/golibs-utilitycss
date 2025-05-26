package types

// Utility represents a container for CSS and Class types.
type Utility interface {
	Css() CSS
	Class() Class
}
