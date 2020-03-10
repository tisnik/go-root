// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
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
