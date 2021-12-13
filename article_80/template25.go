package main

import (
	"os"
	"strings"
	"text/template"
)

const (
	templateValue = `Names: {{join . ", "}}`
)

func main() {
	// mapa funkcí použitých v šabloně
	funcs := template.FuncMap{"join": strings.Join}

	// vytvoření nové šablony
	tmpl := template.Must(template.New("template").Funcs(funcs).Parse(templateValue))

	// tyto hodnoty budou použity při aplikaci šablony
	roles := []string{
		"Eliška Najbrtová",
		"Jenny Suk",
		"Anička Šafářová",
		"Sváťa Pulec",
		"Blažej Motyčka",
		"Eda Wasserfall",
		"Přemysl Hájek",
	}

	// aplikace šablony - přepis hodnot
	err := tmpl.Execute(os.Stdout, roles)
	if err != nil {
		panic(err)
	}
}
