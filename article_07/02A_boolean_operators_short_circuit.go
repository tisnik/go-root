// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_07/README.md
//
// Demonstrační příklad číslo 2:
//    Logické operátory.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/02A_boolean_operators_short_circuit.html

package main

import "fmt"

func f1() bool {
	println("f1")
	return true
}

func f2() bool {
	println("f2")
	return false
}

func f3() bool {
	println("f2")
	return false
}

func main() {
	fmt.Printf("short circuit &&: %v\n", f1() && f2())
	fmt.Printf("short circuit ||: %v\n", f1() || f2())
	fmt.Printf("short circuit &&: %v\n", f2() && f3())
	fmt.Printf("short circuit ||: %v\n", f2() || f3())
}
