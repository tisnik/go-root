// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 13:
//    Defer a příkaz return.

package main

import "fmt"

func function(x int) {
	fmt.Printf("Defer %d\n", x)
}

func classify(x int) string {
	defer function(x)
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
