// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Desátá část
//     Užitečné balíčky pro každodenní použití Go (2), porovnání výkonnosti Go s céčkem
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-go-2-porovnani-vykonnosti-go-s-ceckem/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z desáté části:
//    https://github.com/tisnik/go-root/blob/master/article_10/README.md
//
// Demonstrační příklad číslo 1:
//    Přečtení argumentů předaných na příkazovém řádku.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_10/01_cmdline_params.html

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Arguments: %d\n", len(os.Args))

	for index, element := range os.Args {
		fmt.Printf("Argument #%d = %s\n", index, element)
	}
}
