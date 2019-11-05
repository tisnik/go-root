// Seriál "Programovací jazyk Go"
//
// Sedmnáctá část
//
// Demonstrační příklad číslo 8:
//     Testovaný balíček.

package main

func add(x int32, y int32) int32 {
	return x + y
}

func main() {
	println(add(1, 2))
}
