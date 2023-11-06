// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Sedmdesátá devátá část:
//    Standardní šablonovací systém jazyka Go
//    https://www.root.cz/clanky/standardni-sablonovaci-system-jazyka-go/
//
// Seznam příkladů ze sedmdesáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_79/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_79/template03.html

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
