// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 7:
//    Řídicí konstrukce if a hodnoty, které nejsou typu boolean.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/07_if_statement_bad_type.html

package main

func main() {
	if 1 {
		println("true")
	}

	if 0 {
		println("false")
	}

	if !1 {
		println("false")
	}

	if !0 {
		println("true")
	}

	var b1 int = 1

	if b1 {
		println("true")
	}
	if !b1 {
		println("false")
	}

	b2 := 1

	if b2 {
		println("true")
	}
	if !b2 {
		println("false")
	}
}
