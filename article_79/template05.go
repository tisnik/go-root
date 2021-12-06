package main

import (
	"os"
	"text/template"
)

const (
	templateName   = "test"
	templateFormat = "Hello {{.}}"
)

func main() {
	// vytvoření nové šablony
	tmpl := template.Must(template.New(templateName).Parse(templateFormat))

	// aplikace šablony - přepis hodnot
	err := tmpl.Execute(os.Stdout, "world")
	if err != nil {
		panic(err)
	}
}
