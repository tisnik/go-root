// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvanáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (pokračování)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-pokracovani/
//
// Demonstrační příklad číslo 11:
//     Jednoduchý HTTPS server
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_12/11_https_server.html

package main

import (
	"io"
	"log"
	"net/http"
)

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	io.WriteString(writer, "Hello world!\n")
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	err := http.ListenAndServeTLS(":4443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
