// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá čtvrtá část
//    Programovací jazyk Go pro skalní céčkaře
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
//
// Seznam příkladů ze třicáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_34/README.md
//
// Demonstrační příklad číslo 15:
//    Deklarace jednoduché funkce.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_34/15_simple_function.html
//

package main

import "fmt"

func printHello() {
	fmt.Println("Hello world!")
}

func main() {
	printHello()
}
