// Seriál "Programovací jazyk Go"
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 16:
//    Variadické funkce.

package main

import "fmt"

func f1(msg string) {
	fmt.Printf("%s\n", msg)
}

func f2(parts ...string) {
	for _, val := range parts {
		fmt.Printf("%s ", val)
	}
	fmt.Println()
}

func f3(prefix string, parts ...string) {
	fmt.Println(prefix)
	for _, val := range parts {
		fmt.Printf("%s ", val)
	}
	fmt.Println()
}

func main() {
	f1("Hello")
	f2("Hello", "world", "!")
	f3("Message:", "Hello", "world", "again", "!")
}
