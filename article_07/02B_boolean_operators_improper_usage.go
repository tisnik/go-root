// Seriál "Programovací jazyk Go"
//
// Sedmá část
//
// Demonstrační příklad číslo 2B:
//    Nekorektní použití operátorů na pravdivostní hodnoty.

package main

import "fmt"

func main() {
	x := true
	y := false

	fmt.Printf("^%v == %v\n", x, ^x)
	fmt.Printf("^%v == %v\n", y, ^y)

	fmt.Printf("%v & %v == %v\n", x, y, x&y)
	fmt.Printf("%v | %v == %v\n", x, y, x|y)
	fmt.Printf("%v + %v == %v\n", x, y, x+y)
	fmt.Printf("%v - %v == %v\n", x, y, x-y)
}
