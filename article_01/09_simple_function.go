// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 9:
//    Deklarace uživatelské funkce
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_01/09_simple_function.go

package main

import "fmt"

func printHello() {
	fmt.Println("Hello world!")
}

func main() {
	printHello()
}
