// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá první část
//    Tvorba grafů v jazyce Go: kreslení ve webovém klientu
//    https://www.root.cz/clanky/tvorba-grafu-v-jazyce-go-kresleni-ve-webovem-klientu/

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
	http.Handle("/", http.FileServer(http.Dir("plotly10/")))
	http.HandleFunc("/values", valuesHandler)
	http.ListenAndServe(address, nil)
}

func main() {
	startHttpServer(":8080")
}
