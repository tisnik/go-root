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
// Demonstrační příklad číslo 27:
//    Příkaz break.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/27_break_statement.html

package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
		if i == 5 {
			break
		}
	}
}
