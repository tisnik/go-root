// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 16:
//    Chybný pokus o deklaraci proměnné za příkazem return
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_01/16_return_at_the_end_of_function.html

package main

import "fmt"

func getMessage() (message string) {
	message = "Hello world!"
	return
	x := 42
}

func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	printMessage(getMessage())
}
