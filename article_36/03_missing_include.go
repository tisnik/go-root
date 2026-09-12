// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá šestá část
//    Kooperace mezi kódem psaným v Go a C: cgo
//    https://www.root.cz/clanky/kooperace-mezi-kodem-psanym-v-go-a-c-cgo/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Demonstrační příklad číslo 3:
//    Chybějící poznámka s příkazem include.
//
// Seznam demonstračních příkladů ze třicáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_36/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_36/03_missing_include.html
//

package main

import "C"

func main() {
	C.puts("Hello world!")
}
