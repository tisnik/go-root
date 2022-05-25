// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Seznam příkladů z páté části:
//    https://github.com/tisnik/go-root/blob/master/article_05/README.md
//
// Demonstrační příklad číslo 15:
//    Řídicí konstrukce switch - pracuje jinak než v C!
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/15_no_fallthrough_in_switch.html

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
