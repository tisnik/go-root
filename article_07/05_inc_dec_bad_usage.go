// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 5:
//    Nekorektní použití operátorů ++ a --.

package main

import "fmt"

func main() {
	x := 1
	fmt.Printf("x = %d\n", x)

	fmt.Printf("x = %d\n", x++)

	fmt.Printf("x = %d\n", x--)
}
