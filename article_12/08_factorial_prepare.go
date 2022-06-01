// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvanáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (pokračování)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-pokracovani/
//
// Seznam příkladů z dvanácté části:
//    https://github.com/tisnik/go-root/blob/master/article_12/README.md
//
// Demonstrační příklad číslo 8:
//    Webová aplikace pro výpočet faktoriálu (příprava)
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_12/08_factorial_prepare.html

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// FactorialPageDynContent holds all required information for factorial page template
type FactorialPageDynContent struct {
	Result int
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("factorial.html")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
		return
	}

	dynData := FactorialPageDynContent{1}
	err = t.Execute(writer, dynData)
	if err != nil {
		println("Error executing template")
	}
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.ListenAndServe(":8000", nil)
}
