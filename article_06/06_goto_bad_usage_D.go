// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Seznam příkladů ze šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_06/README.md
//
// Demonstrační příklad číslo 6B:
//    Špatné použití příkazu goto.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/06_goto_bad_usage_D.html

package main

func a() {
	goto FuncB
}

func b() {
FuncB:
}

func main() {
}
