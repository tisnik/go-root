// Seriál "Programovací jazyk Go"
//
// Třicátá pátá část
//
// Demonstrační příklad číslo 1:
//     Pravdivostní hodnoty v jazyce Go.

package main

func main() {
	if nil {
		println("true")
	}

	if !nil {
		println("false")
	}

	if 1 {
		println("true")
	}

	if 0 {
		println("false")
	}

	if 1.5 {
		println("true")
	}

	if 0.0 {
		println("false")
	}

	var b1 bool = nil

	if b1 {
		println("true")
	}

	if !b1 {
		println("false")
	}
}
