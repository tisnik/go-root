// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Osmdesátá část:
//    Standardní šablonovací systém jazyka Go a šablony HTML stránek
//    https://www.root.cz/clanky/standardni-sablonovaci-system-jazyka-go-a-sablony-html-stranek/
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_80/template28.html

package main

import (
	"os"
	"text/template"
)

const (
	templateValue = `{{range .}}{{range .}}{{printf "%3d" .}} {{end}}
{{end}}`
)

func main() {
	const N = 10

	// tabulka s malou násobilkou
	var multiplyTable [N][N]int

	// naplnění tabulky
	for j := 0; j < N; j++ {
		for i := 0; i < N; i++ {
			multiplyTable[j][i] = (i + 1) * (j + 1)
		}
	}

	// vytvoření nové šablony
	tmpl := template.Must(template.New("multiply_table").Parse(templateValue))

	// aplikace šablony - přepis hodnot
	err := tmpl.Execute(os.Stdout, multiplyTable)
	if err != nil {
		panic(err)
	}
}
