// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 9:
//    Řídicí konstrukce if-else.

package main

func classify_char(c rune) string {
	if c >= 'a' && c <= 'z' {
		return "male pismeno"
	} else if c >= 'A' && c <= 'Z' {
		return "velke pismeno"
	} else {
		return "neco jineho"
	}
}

func main() {
	if true {
		println("true")
	} else {
		println("false")
	}

	if !true {
		println("false")
	} else {
		println("true")
	}

	println(classify_char('a'))
	println(classify_char('Z'))
	println(classify_char('?'))
}
