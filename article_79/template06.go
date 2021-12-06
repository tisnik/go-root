package main

import (
	"os"
	"text/template"
)

const (
	templateName   = "test"
	templateFormat = "Součet {{.X}} + {{.Y}} = {{.Z}}"
)

type Expression struct {
	X int
	Y int
	Z int
}

func main() {
	// vytvoření nové šablony
	tmpl := template.Must(template.New(templateName).Parse(templateFormat))

	// tyto hodnoty budou použity při aplikaci šablony
	expression := Expression{
		X: 10,
		Y: 20,
		Z: 30,
	}

	// aplikace šablony - přepis hodnot
	err := tmpl.Execute(os.Stdout, expression)
	if err != nil {
		panic(err)
	}
}
