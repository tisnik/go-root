// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 17:
//    Řídicí konstrukce switch: kombinace podmínek a fallthrough.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/17_switch_combinations.html

package main

func classify(x int) string {
	switch {
	case x == 0:
		return "nula"
	case x%2 == 1:
		return "liché číslo"
	case x%2 == 0:
		fallthrough
	default:
		return "sudé číslo"
	}
}

func main() {
	for x := 0; x <= 10; x++ {
		println(x, classify(x))
	}
}
