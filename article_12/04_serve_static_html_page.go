// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvanáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (pokračování)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-pokracovani/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvanácté části:
//    https://github.com/tisnik/go-root/blob/master/article_12/README.md
//
// Demonstrační příklad číslo 4:
//     Server, který vrací statické HTML stránky načtené ze souborů
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_12/04_serve_static_html_page.html

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendStaticPage(writer http.ResponseWriter, filename string) {
	body, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Fprint(writer, string(body))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
	}
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	sendStaticPage(writer, "index.html")
}

func missingPageEndpoint(writer http.ResponseWriter, request *http.Request) {
	sendStaticPage(writer, "missing.html")
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.HandleFunc("/missing", missingPageEndpoint)
	http.ListenAndServe(":8000", nil)
}
