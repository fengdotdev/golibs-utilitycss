package helpers

import "fmt"

// NewSelectorBlock creates a CSS block with the given selector and CSS properties.
// It returns a string formatted as ".selector { css }".
// Example: NewSelectorBlock("bg-white", "background-color: #ffffff;")
// This is useful for generating utility classes dynamically.
func NewSelectorBlock(selector string, css string) string {

	return fmt.Sprintf(".%s { %s }", selector, css)
}
