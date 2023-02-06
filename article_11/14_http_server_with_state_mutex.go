// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z jedenácté části:
//    https://github.com/tisnik/go-root/blob/master/article_11/README.md
//
// Demonstrační příklad číslo 14:
//     Korektní práce se stavovou proměnnou
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/14_http_server_with_state_mutex.html

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

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.HandleFunc("/counter", counterEndpoint)
	http.ListenAndServe(":8000", nil)
}
