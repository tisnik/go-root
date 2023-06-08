package main

import (
	"html/template"
	"os"
)

const (
	templateFilename = "html_template02.htm"
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
