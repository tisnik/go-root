// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devatenáctá část
//     Využití WebAssembly z programovacího jazyka Go
//     https://www.root.cz/clanky/vyuziti-webassembly-z-programovaciho-jazyka-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z devatenácté části:
//    https://github.com/tisnik/go-root/blob/master/article_19/README.md
//
// Demonstrační příklad:
//    HTTP server vracející statický obsah
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_19/file_server.html

package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("")))
	http.ListenAndServe(":8000", nil)
}
