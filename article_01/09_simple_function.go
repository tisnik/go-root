// Seriál "Programovací jazyk Go"
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 9:
//    Deklarace uživatelské funkce

package main

import "fmt"

func printHello() {
	fmt.Println("Hello world!")
}

func main() {
	printHello()
}
