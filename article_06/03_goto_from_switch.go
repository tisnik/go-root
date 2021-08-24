// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 3:
//    Výskok z příkazu switch pomocí goto.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/03_goto_from_switch.html

package main

func classify(x int) string {
	switch x {
	case 0:
		return "nula"
	case 2, 4, 6, 8:
		goto SudeCislo
	case 1, 3, 5, 7, 9:
		goto LicheCislo
	default:
		goto JineCislo
	}
JineCislo:
	return "?"
SudeCislo:
	return "sudé číslo"
LicheCislo:
	return "liché číslo"
}

func main() {
	for x := 0; x <= 10; x++ {
		println(x, classify(x))
	}
}
