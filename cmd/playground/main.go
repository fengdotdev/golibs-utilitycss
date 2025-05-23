package main

import (
	"fmt"

	. "github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/classes"
	"github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/gorenderercss"
)

func main() {

	renderer := gorenderercss.NewRendererCSS()

	var _ = renderer.Mix(Bg_white, Bg_black, "bg-white", "bg-black")

	class := renderer.Utils(Bg_white, Bg_black)

	fmt.Println("Class: ", class)
	render := renderer.Render()
	fmt.Println("Render: ", render.String())

}
