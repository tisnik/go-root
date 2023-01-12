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
// Demonstrační příklad číslo 10:
//    Příklady špatné syntaxe konstrukce if-else.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/10_bad_syntax.html

package main

func classify_char(c rune) string {
	if c >= 'a' && c <= 'z'
	{
		return "male pismeno"
	}
	else if c >= 'A' && c <= 'Z' {
		return "velke pismeno"
	}
	else
	{
		return "neco jineho"
	}
}

func main() {
	println(classify_char('a'))
	println(classify_char('Z'))
	println(classify_char('?'))
}
