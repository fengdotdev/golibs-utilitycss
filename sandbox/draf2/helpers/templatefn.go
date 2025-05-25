package helpers

import (
	"bytes"
	"fmt"
	htmlTmpl "html/template"
	txtTmpl "text/template"
)

func TemplateGenerator(templateString string, Obj interface{}, isHTML bool) (string, error) {
	if isHTML {
		return HTMLTemplateGenerator(templateString, Obj)
	}

	return TextTemplateGenerator(templateString, Obj)
}

func HTMLTemplateGenerator(templateString string, obj interface{}) (string, error) {

	tmpl, err := htmlTmpl.New("dynamic_template").Parse(templateString)
	if err != nil {
		return "", fmt.Errorf("error al parsear el template: %w", err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, obj)
	if err != nil {
		return "", fmt.Errorf("error al ejecutar el template: %w", err)
	}

	return buf.String(), nil
}

func TextTemplateGenerator(templateString string, obj interface{}) (string, error) {

	tmpl, err := txtTmpl.New("dynamic_template").Parse(templateString)
	if err != nil {
		return "", fmt.Errorf("error al parsear el template: %w", err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, obj)
	if err != nil {
		return "", fmt.Errorf("error al ejecutar el template: %w", err)
	}

	return buf.String(), nil
}
