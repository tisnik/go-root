// Seriál "Programovací jazyk Go"
//
// Dvanáctá část
//
// Demonstrační příklad číslo 4:
//     Server, který vrací statické HTML stránky načtené ze souborů

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
