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
	"log"
	"net/http"
)

func startHttpServer(address string) {
	log.Printf("Starting server on address %s", address)
	http.Handle("/", http.FileServer(http.Dir("plotly01/")))
	http.ListenAndServe(address, nil)
}

func main() {
	startHttpServer(":8080")
}
