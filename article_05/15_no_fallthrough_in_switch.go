// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 15:
//    Řídicí konstrukce switch - pracuje jinak než v C!

package main

func classify(x int) string {
	switch x {
	case 0:
		return "nula"
	case 2:
	case 4:
	case 6:
	case 8:
		return "sudé číslo"
	case 1:
	case 3:
	case 5:
	case 7:
	case 9:
		return "liché číslo"
	default:
		return "?"
	}
	return "X"
}

func main() {
	for x := 0; x <= 10; x++ {
		println(x, classify(x))
	}
}
