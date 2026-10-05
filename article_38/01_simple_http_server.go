// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá osmá část
//    Tvorba webových aplikací v Go s využitím projektu Gorilla web toolkit
//    https://www.root.cz/clanky/tvorba-webovych-aplikaci-v-go-s-vyuzitim-projektu-gorilla-web-toolkit/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Demonstrační příklad číslo 1:
//    Jednoduchý HTTP server založený na standardním balíčku net.
//
// Seznam demonstračních příkladů ze třicáté osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_38/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/html-root/article_38/01_simple_http_server.html
//

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

// ADDRESS represents base address for the HTTP/HTTPs server
const ADDRESS = ":8080"

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

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.HandleFunc("/counter", counterEndpoint)

	log.Println("Starting HTTP server at address", ADDRESS)
	err := http.ListenAndServe(ADDRESS, nil)
	if err != nil {
		log.Fatal("Unable to initialize HTTP server", err)
		os.Exit(2)
	}
}
