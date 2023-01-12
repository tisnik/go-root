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
// Demonstrační příklad číslo 4:
//    Příkaz return ve funkci s návratovou hodnotou.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/04_return_statement_int_value.html

package main

func f2() int {
	println("f2() před příkazem return")
	return 42
	println("f2() po příkazu return")
}

func main() {
	println("Hello world!")
	println(f2())
}
