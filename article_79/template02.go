package main

import (
	"os"
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

	// aplikace šablony - přepis hodnot + výpis výsledku
	err = tmpl.ExecuteTemplate(os.Stdout, templateName, nil)
	if err != nil {
		panic(err)
	}
}
