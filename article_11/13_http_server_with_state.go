// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 13:
//     HTTP server se stavovou proměnnou

package main

import (
	"fmt"
	"io"
	"net/http"
)

var counter int

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!\n")
}

func counterEndpoint(writer http.ResponseWriter, request *http.Request) {
	counter++
	fmt.Fprintf(writer, "Counter: %d\n", counter)
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.HandleFunc("/counter", counterEndpoint)
	http.ListenAndServe(":8000", nil)
}
