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
// Demonstrační příklad číslo 7:
//    Základní použití šablon
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_12/07_html_templates.html

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// IndexPageDynContent holds all required information for a main page template
type IndexPageDynContent struct {
	Title  string
	Header string
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("index_template.html")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
		return
	}

	dynData := IndexPageDynContent{Title: "Test", Header: "Welcome!"}
	err = t.Execute(writer, dynData)
	if err != nil {
		println("Error executing template")
	}
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.ListenAndServe(":8000", nil)
}
