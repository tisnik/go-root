// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 21:
//    Základní forma programové smyčky for s podmínkou.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/21_for_loop_with_condition.html

package main

func main() {
	i := 10
	for i != 0 {
		println(i)
		i--
	}
}
