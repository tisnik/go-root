package main

import (
	"os"
	"text/template"
)

const (
	templateName   = "test"
	templateFormat = `Jméno {{.name}} {{.surname}}
Popularita {{.popularity}}

`
)

type Role struct {
	name       string
	surname    string
	popularity int
}

func main() {
	// vytvoření nové šablony
	tmpl := template.Must(template.New(templateName).Parse(templateFormat))

	// tyto hodnoty budou použity při aplikaci šablony
	roles := []Role{
		Role{"Eliška", "Najbrtová", 4},
		Role{"Jenny", "Suk", 3},
		Role{"Anička", "Šafářová", 1},
		Role{"Sváťa", "Pulec", 3},
		Role{"Blažej", "Motyčka", 8},
		Role{"Eda", "Wasserfall", 3},
		Role{"Přemysl", "Hájek", 10},
	}

	// aplikace šablony - přepis hodnot
	for _, role := range roles {
		err := tmpl.Execute(os.Stdout, role)
		if err != nil {
			panic(err)
		}
	}
}
