// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 4:
//    Nekorektní použití operátorů ++ a --.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/04_inc_dec_bad_usage.html

package main

import "fmt"

func main() {
	x := 1
	fmt.Printf("x = %d\n", x)

	++x
	fmt.Printf("x = %d\n", x)

	--x
	fmt.Printf("x = %d\n", x)
}
