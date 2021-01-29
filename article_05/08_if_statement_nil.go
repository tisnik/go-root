// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 8:
//    Řídicí konstrukce if a hodnoty, které nejsou typu boolean.

package main

func main() {
	if nil {
		println("true")
	}

	if !nil {
		println("false")
	}

	if "" {
		println("true")
	}

	if !"" {
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
