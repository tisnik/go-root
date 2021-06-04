// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá první část
//    Tvorba grafů v jazyce Go: kreslení ve webovém klientu
//    https://www.root.cz/clanky/tvorba-grafu-v-jazyce-go-kresleni-ve-webovem-klientu/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func dataHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, `
{
    "x": [1, 2, 3, 4, 5],
    "y": [1, 2, 4, 8, 16]
}
	`)
}

func startHttpServer(address string) {
	log.Printf("Starting server on address %s", address)
	http.Handle("/", http.FileServer(http.Dir("plotly03/")))
	http.HandleFunc("/data", dataHandler)
	http.ListenAndServe(address, nil)
}

func main() {
	startHttpServer(":8080")
}
