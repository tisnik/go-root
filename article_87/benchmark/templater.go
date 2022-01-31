package main

import (
	"bytes"
	"fmt"
	"html/template"
)

const (
	templateFilename = "html_template05.htm"
)

// datový typ, jehož prvky budou vypisovány v šabloně
type Role struct {
	Name       string
	Surname    string
	Popularity int
}

func readTemplate(templateFilename string) *template.Template {
	// vytvoření nové šablony
	return template.Must(template.ParseFiles(templateFilename))
}

func applyTemplate(tmpl *template.Template, roles []Role) int {
	buffer := new(bytes.Buffer)
	// aplikace šablony - přepis hodnot
	err := tmpl.Execute(buffer, roles)
	if err != nil {
		panic(err)
	}

	return buffer.Len()
}

func main() {
	// tyto hodnoty budou použity při aplikaci šablony
	roles := []Role{
		Role{"Eliška", "Najbrtová", 4},
		Role{"Jenny", "Suk", 3},
		Role{"Anička", "Šafářová", 0},
		Role{"Sváťa", "Pulec", 3},
		Role{"Blažej", "Motyčka", 8},
		Role{"Eda", "Wasserfall", 0},
		Role{"Přemysl", "Hájek", 10},
	}

	tmpl := readTemplate(templateFilename)

	fmt.Println(applyTemplate(tmpl, roles))
}
