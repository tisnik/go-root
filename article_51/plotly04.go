// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá první část
//    Tvorba grafů v jazyce Go: kreslení ve webovém klientu
//    https://www.root.cz/clanky/tvorba-grafu-v-jazyce-go-kresleni-ve-webovem-klientu/
//
// Seznam příkladů z padesáté první části:
//    https://github.com/tisnik/go-root/blob/master/article_51/README.md

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Points struct {
	X []int `json:"x"`
	Y []int `json:"y"`
}

func dataHandler(writer http.ResponseWriter, request *http.Request) {
	var points Points
	points.X = []int{1, 2, 3, 4, 5}
	points.Y = []int{1, 2, 4, 8, 16}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(points)

}

func startHttpServer(address string) {
	log.Printf("Starting server on address %s", address)
	http.Handle("/", http.FileServer(http.Dir("plotly04/")))
	http.HandleFunc("/data", dataHandler)
	http.ListenAndServe(address, nil)
}

func main() {
	startHttpServer(":8080")
}
