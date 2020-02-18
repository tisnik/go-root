// Seriál "Programovací jazyk Go"
//
// Dvanáctá část
//
// Demonstrační příklad číslo 7:
//    Základní použití šablon

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
