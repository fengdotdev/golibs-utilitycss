package main

import (
	"fmt"

	"github.com/fengdotdev/golibs-utilitycss/v0/helpers"
)

var htmlContent = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Ejemplo</title>
	</head>
	<body>
		<div class="header main-class"></div>
		<p class="content-text"></p>
		<span id="item" class='some-item another-class'></span>
		<a href="#" class="link"></a>
		<div class="footer"></div>
		<div>Sin clase</div>
		<img src="image.png" class="image-style" alt="Una imagen">
		<input type="text" value="algo" class ="input-field" />
	</body>
	</html>
	`

func main() {
	url := "https://preline.co/examples.html"
	htmlremote, err := helpers.FetchHTML(url)

	if err != nil {
		panic(err)
	}
	set, err := helpers.Parser(htmlremote)
	helpers.ForEach(set, func(class string) {
		s := fmt.Sprintf("at %s class: %s", url, class)
		fmt.Println(s)
	})

}
