package helpers

import (
	"strings"

	"golang.org/x/net/html"
)

func FinderClasses(n *html.Node, finds chan string) {

	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "class" {

				// Split the class attribute value into individual classes
				classes := strings.Fields(a.Val)
				for _, class := range classes {
					// Send each class to the channel
					finds <- class
				}

				break // Stop searching after the first class attribute
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		FinderClasses(c, finds)
	}
}
