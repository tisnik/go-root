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
// Demonstrační příklad číslo 13:
//    Řídicí konstrukce switch s vyhodnocovaným výrazem.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/13_variables_in_switch.html

package main

func classify(x int, zero_value int) string {
	switch x {
	case zero_value:
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
		println(x, classify(x, 0))
	}
}
