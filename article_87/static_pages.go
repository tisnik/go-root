// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů z osmdesáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_87/README.md

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func sendStaticPage(writer http.ResponseWriter, filename string) {
	log.Printf("Sending static file %s", filename)

	body, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Fprint(writer, string(body))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
	}
}

func staticPage(filename string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		sendStaticPage(writer, filename)
	}
}

func main() {
	const address = ":8080"

	log.Printf("Starting server on address %s", address)

	http.HandleFunc("/", staticPage("index.html"))
	http.HandleFunc("/missing", staticPage("missing.html"))
	http.ListenAndServe(address, nil)
}
