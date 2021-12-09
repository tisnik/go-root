// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá devátá část:
//    Standardní šablonovací systém jazyka Go
//    https://www.root.cz/clanky/standardni-sablonovaci-system-jazyka-go/
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_79/template02.html

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
