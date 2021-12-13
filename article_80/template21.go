package main

import (
	"fmt"
	"os"
	"text/template"
)

const (
	templateFilename = "template21.txt"
)

// datový typ, jehož prvky budou vypisovány v šabloně
type Role struct {
	Name       string
	Surname    string
	Popularity int
}

func (role Role) GetPopularity() string {
	if role.Popularity <= 0 {
		return "Nezadáno"
	} else {
		return fmt.Sprintf("%d", role.Popularity)
	}
}

func main() {
	// vytvoření nové šablony
	tmpl := template.Must(template.ParseFiles(templateFilename))

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

	// aplikace šablony - přepis hodnot
	err := tmpl.Execute(os.Stdout, roles)
	if err != nil {
		panic(err)
	}
}
