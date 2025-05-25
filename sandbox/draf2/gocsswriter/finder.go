package gocsswriter

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Set = map[string]struct{}
type List = []string

func Parser(htmlContent string) List {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Printf("Error al parsear HTML: %v\n", err)
		return nil
	}

	// search for classes

	finds := make(chan string)

	go func() {
		FindClasses(doc, finds)
		close(finds)

	}()

	classes := make(Set)
	for class := range finds {
		// Add each class to the set
		classes[class] = struct{}{}
	}
	// Convert the set to a list
	list := SetToList(classes)
	return list
}

func FindClasses(n *html.Node, finds chan string) {

	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "class" {
				fmt.Printf("Found class: %s\n", a.Val)

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
		FindClasses(c, finds)
	}
}

func SetToList(set Set) List {
	list := make([]string, 0, len(set))
	for k := range set {
		list = append(list, k)
	}
	return list
}

type Mapeable interface {
	map[string]struct{} | []string
}

func ForEach[T Mapeable](Mapeable T, f func(string)) error {
	switch v := any(Mapeable).(type) {
	case Set:
		ForEachSet(v, f)
		return nil
	case List:
		ForEachList(v, f)
		return nil
	default:
		return errors.New("unsupported type")
	}
}

func ForEachSet(set Set, f func(string)) {
	for k := range set {
		f(k)
	}
}

func ForEachList(list List, f func(string)) {
	for _, k := range list {
		f(k)
	}
}
