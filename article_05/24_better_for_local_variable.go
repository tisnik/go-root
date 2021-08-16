// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 24:
//    Zjednodušená forma programové smyčky for odvozené od C.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/24_better_for_local_variable.html

package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}
	println(i)
}
