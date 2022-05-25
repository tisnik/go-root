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
// Příkaz zapsaný za klíčovým slovem if.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/10B_statement_in_if.html

package main

func funkce() int {
	return -1
}

func x() string {
	if value := funkce(); value < 0 {
		return "záporná hodnota"
	} else if value > 0 {
		return "kladná hodnota"
	} else {
		return "nula"
	}
}

func main() {
	println(x())
}
