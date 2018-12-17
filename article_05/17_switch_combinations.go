// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 17:
//    Řídicí konstrukce switch: kombinace podmínek a fallthrough.

package main

func classify(x int) string {
	switch {
	case x == 0:
		return "nula"
	case x%2 == 1:
		return "liché číslo"
	case x%2 == 0:
		fallthrough
	default:
		return "sudé číslo"
	}
}

func main() {
	for x := 0; x <= 10; x++ {
		println(x, classify(x))
	}
}
