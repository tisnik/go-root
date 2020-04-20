// Seriál "Programovací jazyk Go"
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Demonstrační příklad číslo 1:
//     Testovaný balíček.

package main

func Add(x int, y int) int {
	return x + y
}

func main() {
	println(Add(1, 2))
}
