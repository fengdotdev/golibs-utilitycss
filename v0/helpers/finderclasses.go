package helpers

import (
	"strings"

	"golang.org/x/net/html"
)

// FinderClasses traverses an HTML node tree and finds all class attributes.
// It sends each class found to the provided channel.
// the class are split into individual classes and processed.
// uses the "golang.org/x/net/html" package to parse HTML nodes.
// ex:   class"some-class another-class" -> some-class, another-class
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

// FinderClassesWithCallback traverses an HTML node tree and finds all class attributes.
// the class are split into individual classes and processed.
// It calls the provided callback function for each class found.
// This function is useful for processing classes without using channels.
// ex:   class"some-class another-class" -> some-class, another-class
func FinderClassesWithCallback(n *html.Node, callback func(string)) {

	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "class" {

				// Split the class attribute value into individual classes
				classes := strings.Fields(a.Val)
				for _, class := range classes {
					// Call the callback function with each class
					callback(class)
				}

				break // Stop searching after the first class attribute
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		FinderClassesWithCallback(c, callback)
	}
}
