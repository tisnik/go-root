package main

import (
	"log"
	"net/http"
)

func startHttpServer(address string) {
	log.Printf("Starting server on address %s", address)
	http.Handle("/", http.FileServer(http.Dir("plotly02/")))
	http.ListenAndServe(address, nil)
}

func main() {
	startHttpServer(":8080")
}
