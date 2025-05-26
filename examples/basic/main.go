package main

import (
	"fmt"

	. "github.com/fengdotdev/golibs-utilitycss/v0/classes"

	"github.com/fengdotdev/golibs-utilitycss/v0/gorenderercss"
)

// Example of using the gorenderercss package to create a CSS class and render it
// this prints the class and the rendered CSS

func main() {
	renderer := gorenderercss.NewRendererCSS()

	class := renderer.Utils(Bg_white, Bg_black)

	fmt.Println("Class: ", class)
	render := renderer.Render()
	fmt.Println("Render: ", render.String())
}
