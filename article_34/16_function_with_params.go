// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//    Programovací jazyk Go pro skalní céčkaře
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
//
// Demonstrační příklad číslo 16:
//    Funkce s parametry.

package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	printMessage("Hello world!")
}
