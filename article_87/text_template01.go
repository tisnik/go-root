// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů z osmdesáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_87/README.md

package main

import (
	"os"
	"text/template"
)

const (
	templateFilename = "text_template01.txt"
)

// datový typ, jehož prvky budou vypisovány v šabloně
type Role struct {
	Name       string
	Surname    string
	Popularity int
}

func main() {
	// vytvoření nové šablony
	tmpl := template.Must(template.ParseFiles(templateFilename))

	// tyto hodnoty budou použity při aplikaci šablony
	roles := []Role{
		{"Eliška", "Najbrtová", 4},
		{"Jenny", "Suk", 3},
		{"Anička", "Šafářová", 0},
		{"Sváťa", "Pulec", 3},
		{"Blažej", "Motyčka", 8},
		{"Eda", "Wasserfall", 0},
		{"Přemysl", "Hájek", 10},
	}

	// aplikace šablony - přepis hodnot
	err := tmpl.Execute(os.Stdout, roles)
	if err != nil {
		panic(err)
	}
}
