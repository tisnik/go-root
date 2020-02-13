// Seriál "Programovací jazyk Go"
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 17:
//    Variadické funkce.

package main

import "fmt"

func f4(prefix string, parts1 ...string, parts2 ...string) {
	fmt.Println(prefix)
	for _, val := range parts1 {
		fmt.Printf("%s ", val)
	}
	fmt.Println()
	for _, val := range parts2 {
		fmt.Printf("%s ", val)
	}
	fmt.Println()
}

func main() {
	f4("Message:", "Hello", "world", "again", "!")
}
