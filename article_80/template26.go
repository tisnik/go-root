// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmdesátá část:
//    Standardní šablonovací systém jazyka Go a šablony HTML stránek
//    https://www.root.cz/clanky/standardni-sablonovaci-system-jazyka-go-a-sablony-html-stranek/
//
// Seznam příkladů z osmdesáté části:
//    https://github.com/tisnik/go-root/blob/master/article_80/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_80/template26.html

package main

import (
	"os"
	"strings"
	"text/template"
)

const (
	templateValue = `Names: {{with $names := asNames .}}{{join $names ", "}}{{end}}`
)

// datový typ, jehož prvky budou vypisovány v šabloně
type Role struct {
	Name       string
	Surname    string
	Popularity int
}

// převod rolí na řez se jmény rolí
func asNames(roles []Role) []string {
	var r []string
	for _, role := range roles {
		r = append(r, role.Name)
	}
	return r
}

func main() {
	// mapa funkcí použitých v šabloně
	funcs := template.FuncMap{
		"asNames": asNames,
		"join":    strings.Join}

	// vytvoření nové šablony
	tmpl := template.Must(template.New("template").Funcs(funcs).Parse(templateValue))

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
