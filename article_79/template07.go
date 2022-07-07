// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá devátá část:
//    Standardní šablonovací systém jazyka Go
//    https://www.root.cz/clanky/standardni-sablonovaci-system-jazyka-go/
//
// Seznam příkladů ze sedmdesáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_79/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_79/template07.html

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
