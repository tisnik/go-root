// Seriál "Programovací jazyk Go"
//
// Sedmá část
//
// Demonstrační příklad číslo 1:
//    Bitové operátory.

package main

import "fmt"

func main() {
	x := 1
	y := 0xfe

	fmt.Printf("%x & %x == %x\n", x, y, x&y)
	fmt.Printf("%x &^ %x == %x\n", x, y, x&^y)
	fmt.Printf("%x | %x == %x\n", x, y, x|y)
	fmt.Printf("%x ^ %x == %x\n", x, y, x^y)

	x ^= y
	fmt.Printf("new x = %x\n", x)

	x |= y
	fmt.Printf("new x = %x\n", x)

	x ^= y
	fmt.Printf("new x = %x\n", x)

	x &^= 0x01
	fmt.Printf("new x = %x\n", x)

	fmt.Println()

	x = 1
	y = 2

	fmt.Printf("%x & %x == %x\n", x, y, x&y)
	fmt.Printf("%x &^ %x == %x\n", x, y, x&^y)
	fmt.Printf("%x | %x == %x\n", x, y, x|y)
	fmt.Printf("%x ^ %x == %x\n", x, y, x^y)

}
