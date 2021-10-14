// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 16:
//     Kombinace předchozích možností
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/16_custom_server.html

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!\n")
}

func counterEndpoint(writer http.ResponseWriter, request *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(writer, "Counter: %d\n", counter)
	mutex.Unlock()
}

func filesEndpoint(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.Path[len("/files/"):]
	println("Serving file from URL: " + url)
	http.ServeFile(writer, request, url)
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.HandleFunc("/counter", counterEndpoint)
	http.HandleFunc("/files/", filesEndpoint)
	http.ListenAndServe(":8000", nil)
}
