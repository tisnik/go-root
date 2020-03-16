// Seriál "Programovací jazyk Go"
//
// Dvanáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (pokračování)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-pokracovani/
//
// Demonstrační příklad číslo 2:
//    Alternativní způsob vytvoření sekvence bajtů z řetězce

package main

import (
	"net/http"
)

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	response := "Hello world!\n"
	writer.Write([]byte(response))
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.ListenAndServe(":8000", nil)
}
