// Seriál "Programovací jazyk Go"
//
// Dvanáctá část
//
// Demonstrační příklad číslo 3:
//     HTTP server se specifikací MIME typu odpovědí

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
