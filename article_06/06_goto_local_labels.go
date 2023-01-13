// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_06/README.md
//
// Demonstrační příklad číslo 6C:
//    Lokální návěští.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/06_goto_local_labels.html

package main

func a() {
	goto Label
Label:
}

func b() {
	goto Label
Label:
}

func main() {
}
