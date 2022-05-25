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
// Demonstrační příklad číslo 16:
//    Řídicí konstrukce switch - simulace chování C.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/16_fallthrough_in_switch.html

package main

func classify(x int) string {
	switch x {
	case 0:
		return "nula"
	case 2:
		fallthrough
	case 4:
		fallthrough
	case 6:
		fallthrough
	case 8:
		return "sudé číslo"
	case 1:
		fallthrough
	case 3:
		fallthrough
	case 5:
		fallthrough
	case 7:
		fallthrough
	case 9:
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
