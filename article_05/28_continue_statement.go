// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 28:
//    Příkaz continue.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/28_continue_statement.html

package main

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		println(i)
	}
}
