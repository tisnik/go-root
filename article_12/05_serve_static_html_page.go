// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvanáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (pokračování)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-pokracovani/
//
// Demonstrační příklad číslo 5:
//     Úprava předchozího příkladu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_12/05_serve_static_html_page.html

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
