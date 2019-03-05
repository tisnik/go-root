// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//
// Demonstrační příklad číslo 15:
//     HTTP server vracející statický obsah

package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("")))
	http.ListenAndServe(":8000", nil)
}
