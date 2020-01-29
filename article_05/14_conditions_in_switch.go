// Seriál "Programovací jazyk Go"
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 14:
//    Řídicí konstrukce switch s výrazy ve větvích case.

package main

func classify(x int) string {
	switch {
	case x == 0:
		return "nula"
	case x%2 == 0:
		return "sudé číslo"
	case x%2 == 1:
		return "liché číslo"
	default:
		return "?"
	}
}

func main() {
	for x := 0; x <= 10; x++ {
		println(x, classify(x))
	}
}
