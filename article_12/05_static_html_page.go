// Seriál "Programovací jazyk Go"
//
// Dvanáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (pokračování)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-pokracovani/
//
// Demonstrační příklad číslo 5:
//    Úprava předchozího příkladu

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadFile("index.html")
	if err == nil {
		fmt.Fprint(writer, string(body))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
	}
}

func missingEndpoint(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadFile("missing.html")
	if err == nil {
		fmt.Fprint(writer, string(body))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
	}
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.HandleFunc("/missing", missingPageEndpoint)
	http.ListenAndServe(":8000", nil)
}
