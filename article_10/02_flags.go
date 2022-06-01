// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Desátá část
//     Užitečné balíčky pro každodenní použití Go (2), porovnání výkonnosti Go s céčkem
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-go-2-porovnani-vykonnosti-go-s-ceckem/
//
// Seznam příkladů z desáté části:
//    https://github.com/tisnik/go-root/blob/master/article_10/README.md
//
// Demonstrační příklad číslo 2:
//    Vylepšené přečtení argumentů předaných na příkazovém řádku.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_10/02_flags.html

package main

import (
	"flag"
	"fmt"
)

func main() {
	var width int
	var height int
	var aa bool
	var output string

	flag.IntVar(&width, "width", 0, "image width")
	flag.IntVar(&height, "height", 0, "image height")
	flag.BoolVar(&aa, "aa", false, "enable antialiasing")
	flag.StringVar(&output, "output", "", "output file name")

	flag.Parse()

	fmt.Printf("width: %d\n", width)
	fmt.Printf("height: %d\n", height)
	fmt.Printf("antialiasing: %t\n", aa)
	fmt.Printf("output file name: %s\n", output)
}
