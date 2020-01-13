package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func valuesHandler(writer http.ResponseWriter, request *http.Request) {
	values := []float64{20, 30, 40, 50}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(values)
}

func startHttpServer(address string) {
	log.Printf("Starting server on address %s", address)
	http.Handle("/", http.FileServer(http.Dir("plotly09/")))
	http.HandleFunc("/values", valuesHandler)
	http.ListenAndServe(address, nil)
}

func main() {
	startHttpServer(":8080")
}
