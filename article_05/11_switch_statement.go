// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 11:
//    Řídicí konstrukce switch bez vyhodnocovaného výrazu.

package main

func main() {
	switch {
	}

	switch {
	default:
		println("proč jsem vlastně použil switch?")
	}

	switch {
	case true:
		println("true")
	case false:
		println("false")
	}

	switch {
	case false:
		println("false")
	case true:
		println("true")
	default:
		println("default")
	}

	switch {
	case false:
		println("false")
	default:
		println("default")
	case true:
		println("true")
	}
}
