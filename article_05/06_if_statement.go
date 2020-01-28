// Seriál "Programovací jazyk Go"
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 6:
//    Řídicí konstrukce if.

package main

func main() {
	if true {
		println("true")
	}

	if false {
		println("false")
	}

	if !true {
		println("false")
	}

	if !false {
		println("true")
	}

	var b1 bool = true

	if b1 {
		println("true")
	}
	if !b1 {
		println("false")
	}

	b2 := true

	if b2 {
		println("true")
	}
	if !b2 {
		println("false")
	}
}
