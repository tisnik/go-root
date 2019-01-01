// Seriál "Programovací jazyk Go"
//
// Sedmá část
//
// Demonstrační příklad číslo 4:
//    Nekorektní použití operátorů ++ a --.

package main

import "fmt"

func main() {
	x := 1
	fmt.Printf("x = %d\n", x)

	++x
	fmt.Printf("x = %d\n", x)

	--x
	fmt.Printf("x = %d\n", x)
}
