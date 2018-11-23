// Seriál "Programovací jazyk Go"
//
// Druhá část
//
// Demonstrační příklad číslo 4:
//    Celočíselné konstanty reprezentované v různých číselných soustavách

package main

import "fmt"

func main() {
	var a uint8 = 10
	var b uint8 = 010
	var c uint8 = 0x10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var d int8 = -10
	var e int8 = -010
	var f int8 = -0x10

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
