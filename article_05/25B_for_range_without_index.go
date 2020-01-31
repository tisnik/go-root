// Seriál "Programovací jazyk Go"
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 25:
//    Varianta s nepoužitým indexem.

package main

func main() {
	a := [...]int{1, 2, 10, -1, 42}

	for i, item := range a {
		println(item)
	}

	println()

	s := "Hello world ěščř Σ"

	for i, character := range s {
		println(character)
	}
}
