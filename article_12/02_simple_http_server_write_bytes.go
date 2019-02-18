package main

import (
	"net/http"
)

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	response := "Hello world!\n"
	writer.Write([]byte(response))
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.ListenAndServe(":8000", nil)
}
