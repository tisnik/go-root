// Seriál "Programovací jazyk Go"
//
// Sedmá část
//
// Demonstrační příklad číslo 3:
//    Operátory ++ a --.

package main

import "fmt"

func main() {
	x := 1
	fmt.Printf("x = %d\n", x)

	x++
	fmt.Printf("x = %d\n", x)

	x--
	fmt.Printf("x = %d\n", x)

	px := &x
	fmt.Printf("x = %d\n", *px)

	*px++
	fmt.Printf("x = %d\n", *px)

	*px--
	fmt.Printf("x = %d\n", *px)
}
