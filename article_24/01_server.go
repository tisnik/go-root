// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá čtvrtá část
//     Kontejnery v základní knihovně programovacího jazyka Go
//     https://www.root.cz/clanky/kontejnery-v-zakladni-knihovne-programovaciho-jazyka-go/
//
// Seznam příkladů z dvacáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_24/README.md
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
