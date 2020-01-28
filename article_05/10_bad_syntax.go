// Seriál "Programovací jazyk Go"
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 10:
//    Příklady špatné syntaxe konstrukce if-else.

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
