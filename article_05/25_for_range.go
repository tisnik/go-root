// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 25:
//    For s range.

package main

func main() {
	a := [...]int{1, 2, 10, -1, 42}

	for index, item := range a {
		println(index, item)
	}

	println()

	s := "Hello world ěščř Σ"

	for index, character := range s {
		println(index, character)
	}
}
