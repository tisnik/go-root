// Seriál "Programovací jazyk Go"
//
// Druhá část
//
// Demonstrační příklad číslo 5:
//    Přiřazení uint8 -> uint16

package main

import "fmt"

func main() {
	var a uint8 = 255
	var b uint16 = a

	fmt.Println(a)
	fmt.Println(b)
}
