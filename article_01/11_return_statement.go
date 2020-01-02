// Seriál "Programovací jazyk Go"
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 11:
//    Deklarace uživatelské funkce s návratovou hodnotou

package main

import "fmt"

func getMessage() string {
	return "Hello world!"
}

func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	printMessage(getMessage())
}
