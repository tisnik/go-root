// Seriál "Programovací jazyk Go"
//
// Dvacátá čtvrtá část
//     Kontejnery v základní knihovně programovacího jazyka Go
//     https://www.root.cz/clanky/kontejnery-v-zakladni-knihovne-programovaciho-jazyka-go/
//
// Demonstrační příklad číslo 1:
//     	Jednoduchý HTTP server posílající dynamicky generovaný obsah

package main

import (
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	println("handler started")
	for i := 0; i < 50; i++ {
		writer.Write([]byte("Hello world!\n"))
	}
	println("handler finished")
}

func main() {
	http.HandleFunc("/data", handler)
	http.ListenAndServe(":8080", nil)
}
