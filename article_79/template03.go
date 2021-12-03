package main

import (
	"bytes"
	"fmt"
	"text/template"
)

const (
	templateName   = "test"
	templateFormat = "Toto je testovací šablona"
)

func main() {
	// vytvoření nové šablony
	tmpl, err := template.New(templateName).Parse(templateFormat)
	if err != nil {
		panic(err)
	}

	// buffer pro uložení výsledků aplikace šablony
	buffer := new(bytes.Buffer)

	// aplikace šablony - přepis hodnot
	err = tmpl.Execute(buffer, nil)
	if err != nil {
		panic(err)
	}

	// výpis výsledného textu
	fmt.Println(buffer.String())
}
