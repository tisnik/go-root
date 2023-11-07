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
//    https://tisnik.github.io/go-root/article_79/template08.html

package main

import (
	"os"
	"text/template"
)

const (
	templateName   = "test"
	templateFormat = "Součet {{.X}} + {{.Y}} = {{.Z}}"
)

type User struct {
	FirstName string
	Surname   string
	Born      string
}

func main() {
	// vytvoření nové šablony
	tmpl := template.Must(template.New(templateName).Parse(templateFormat))

	// tyto hodnoty budou použity při aplikaci šablony
	user := User{
		FirstName: "Jára",
		Surname:   "Cimrman",
		Born:      "Böhmen",
	}

	// aplikace šablony - přepis hodnot
	err := tmpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
