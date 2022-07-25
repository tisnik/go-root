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
//    https://tisnik.github.io/go-root/article_80/template29.html

package main

import (
	"os"
	"text/template"
)

const (
	templateValue = `{{range .Values}}{{range .}}{{printf "%3d" .}} {{end}}
{{end}}`
)

const N = 10

type MultiplyTable struct {
	Values [N][N]int
}

func main() {

	// tabulka s malou násobilkou
	var multiplyTable MultiplyTable

	// naplnění tabulky
	for j := 0; j < N; j++ {
		for i := 0; i < N; i++ {
			multiplyTable.Values[j][i] = (i + 1) * (j + 1)
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
