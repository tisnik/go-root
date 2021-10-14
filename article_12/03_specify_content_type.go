// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvanáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (pokračování)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-pokracovani/
//
// Demonstrační příklad číslo 3:
//     HTTP server se specifikací MIME typu odpovědí
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_12/03_specify_content_type.html

package main

import (
	"net/http"
)

func endpointHTML(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	response := "<body><h1>Hello world!</h1></body>\n"
	writer.Write([]byte(response))
}

func endpointText(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	response := "Hello world!\n"
	writer.Write([]byte(response))
}

func endpointAsm(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/x-asm")
	response := "START: brk\n"
	writer.Write([]byte(response))
}

func main() {
	http.HandleFunc("/html", endpointHTML)
	http.HandleFunc("/text", endpointText)
	http.HandleFunc("/asm", endpointAsm)
	http.ListenAndServe(":8000", nil)
}
