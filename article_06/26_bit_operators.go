// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 26:
//    Bitové operátory.

package main

import "fmt"

func main() {
	x := 1
	y := 0xfe

	fmt.Printf("%x & %x == %x\n", x, y, x&y)
	fmt.Printf("%x | %x == %x\n", x, y, x|y)
	fmt.Printf("%x ^ %x == %x\n", x, y, x^y)

	x ^= y
	fmt.Printf("new x = %x\n", x)

	x |= y
	fmt.Printf("new x = %x\n", x)

	x ^= y
	fmt.Printf("new x = %x\n", x)
}
