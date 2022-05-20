// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmdesátá část:
//    Standardní šablonovací systém jazyka Go a šablony HTML stránek
//    https://www.root.cz/clanky/standardni-sablonovaci-system-jazyka-go-a-sablony-html-stranek/
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_80/template24.html

package main

import (
	"os"
	"text/template"
)

const (
	templateValue = `Roles:{{block "roles" .}}{{"\n"}}{{range .}}{{println "-" .Name "\t" .Surname}}{{end}}{{end}}`
)

// datový typ, jehož prvky budou vypisovány v šabloně
type Role struct {
	Name       string
	Surname    string
	Popularity int
}

func main() {
	// vytvoření nové šablony
	tmpl := template.Must(template.New("template").Parse(templateValue))

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
