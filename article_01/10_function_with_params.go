// Seriál "Programovací jazyk Go"
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
