// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Seznam příkladů ze sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_07/README.md
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
