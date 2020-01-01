// Seriál "Programovací jazyk Go"
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 10:
//    Deklarace uživatelské funkce s jedním parametrem

package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	printMessage("Hello world!")
}
