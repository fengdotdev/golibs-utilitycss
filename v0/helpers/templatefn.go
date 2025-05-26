package helpers

import (
	"bytes"
	"fmt"
	htmlTmpl "html/template"
	txtTmpl "text/template"
)


// TemplateGenerator generates a template based on the provided template string and object.
// It supports both HTML and text templates based on the isHTML flag.
// returns the rendered template as a string or an error if the rendering fails.
// It uses the html/template package for HTML templates and text/template for text templates.
func TemplateGenerator(templateString string, Obj interface{}, isHTML bool) (string, error) {
	if isHTML {
		return HTMLTemplateGenerator(templateString, Obj)
	}

	return TextTemplateGenerator(templateString, Obj)
}


// HTMLTemplateGenerator generates an HTML template based on the provided template string and object.
// It uses the html/template package to parse and execute the template.
func HTMLTemplateGenerator(templateString string, obj interface{}) (string, error) {

	tmpl, err := htmlTmpl.New("dynamic_template").Parse(templateString)
	if err != nil {
		return "", fmt.Errorf("error parsing template: %w", err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, obj)
	if err != nil {
		return "", fmt.Errorf(" error to execute template: %w", err)
	}

	return buf.String(), nil
}


// TextTemplateGenerator generates a text template based on the provided template string and object.
// It uses the text/template package to parse and execute the template.
func TextTemplateGenerator(templateString string, obj interface{}) (string, error) {

	tmpl, err := txtTmpl.New("dynamic_template").Parse(templateString)
	if err != nil {
		return "", fmt.Errorf("error parsing template: %w", err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, obj)
	if err != nil {
		return "", fmt.Errorf("error to execute on template: %w", err)
	}

	return buf.String(), nil
}
