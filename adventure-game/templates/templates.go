package templates

import (
	"fmt"
	"html/template"
)

func CreateTemplate(fileName string) (*template.Template, error) {
	file := fileName + ".html"
	t, err := template.ParseFiles(file)

	if err != nil {
		return nil, fmt.Errorf("there is some error templating the file %w", err)
	}

	return t, nil
}
