// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//    Programovací jazyk Go pro skalní céčkaře
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
//
// Demonstrační příklad číslo 17:
//    Příkaz return pro ukončení funkce.

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
