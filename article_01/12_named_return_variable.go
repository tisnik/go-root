// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk:
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 12:
//    Použití pojmenovaného návratového parametru
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_01/12_named_return_variable.html

package main

import "fmt"

func getMessage() (message string) {
	message = "Hello world!"
	return
}

func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	printMessage(getMessage())
}
