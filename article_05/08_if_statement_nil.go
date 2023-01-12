// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z páté části:
//    https://github.com/tisnik/go-root/blob/master/article_05/README.md
//
// Demonstrační příklad číslo 8:
//    Řídicí konstrukce if a hodnoty, které nejsou typu boolean.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/08_if_statement_nil.html

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
