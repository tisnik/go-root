// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 10:
//    Deklarace uživatelské funkce s jedním parametrem
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_01/10_function_with_params.html

package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	printMessage("Hello world!")
}
