package demo

import (
	"errors"
	"io"
	"strings"
)

var pageTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>{{.Title}}</title>
	{{.RefCSS}}
</head>
<body>
	{{.Body}}
</body>
</html>
`
// PageDTO is the data transfer object for PageRender.
type PageDTO struct {
	Title  string
	Body   string
	RefCSS string
}



// Page creates a new PageRender with the given title and children components.
// If no children are provided, an empty page will be created.
func Page(title string, children ...Component) PageRender {

	return PageRender{
		title:    title,
		children: children,
	}
}



// PageRender is a component that represents a complete HTML page.
type PageRender struct {
	title           string
	html            string // need build
	css             string // need build
	children        []Component
	haveCss         bool
	haveInternalCss bool
	haveExternalCss bool
	built           bool
}

// Build builds the page by rendering the HTML and CSS.
func (p *PageRender) Build() error {
	if p.built {
		return errors.New("page already built")
	}

	var sb strings.Builder
	sb.WriteString(p.html)
	for _, child := range p.children {
		content, err := child.Render()
		if err != nil {
			return err
		}
		sb.WriteString(content)
	}

	p.html = sb.String()
	p.built = true
	return nil
}

// returns the rendered HTML of the page
func (p *PageRender) BuildToString() string {
	if !p.built {
		p.Build()
	}
	return p.html
}


// BuildToWriter writes the rendered HTML of the page to the given io.Writer.
func (p *PageRender) BuildToWriter(wr io.Writer) error {
	if !p.built {
		p.Build()
	}
	_, err := wr.Write([]byte(p.html))
	return err
}


// BuildToSingleFile builds the page to a single file, including both HTML and CSS.
func (p *PageRender) BuildToSingleFile(path string) error {
	panic("unimplemented")
}

// build the css and the html files
func (p *PageRender) BuildTOFiles(path string) error {
	panic("unimplemented")
}

// FromDTO implements trait.DataTransferObject.
func (p *PageRender) ToDTO() (PageDTO, error) {
	dto := PageDTO{
		Title:  p.title,
		Body:   "", // TODO
		RefCSS: p.css,
	}
	return dto, nil
}
