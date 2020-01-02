// Seriál "Programovací jazyk Go"
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk:
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 12:
//    Použití pojmenovaného návratového parametru

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
