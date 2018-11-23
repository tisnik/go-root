// Seriál "Programovací jazyk Go"
//
// Druhá část
//
// Demonstrační příklad číslo 1:
//    Celočíselné datové typy se znaménkem

package main

import "fmt"

func main() {
	var a int8 = -10
	var b int16 = -1000
	var c int32 = -10000
	var d int32 = -1000000

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
