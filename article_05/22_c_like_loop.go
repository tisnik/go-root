// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 22:
//    Programová smyčka for odvozená od C.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/22_c_like_loop.html

package main

func main() {
	var i int
	for i = 0; i < 10; i++ {
		println(i)
	}
	println()
	println(i)
}
