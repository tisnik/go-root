// Seriál "Programovací jazyk Go"
//
// Druhá část
//
// Demonstrační příklad číslo 6:
//    Přiřazení uint16 -> uint8

package main

import "fmt"

func main() {
	var a uint16 = 255
	var b uint8 = a

	fmt.Println(a)
	fmt.Println(b)
}
