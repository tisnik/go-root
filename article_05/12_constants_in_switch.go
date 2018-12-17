// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 12:
//    Řídicí konstrukce switch s vyhodnocovaným výrazem.

package main

func classify(x int) string {
	switch x {
	case 0:
		return "nula"
	case 2, 4, 6, 8:
		return "sudé číslo"
	case 1, 3, 5, 7, 9:
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
