// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Demonstrační příklad číslo 1:
//     Testovaný balíček.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_17/test01/add.html

package main

func Add(x int, y int) int {
	return x + y
}

func main() {
	println(Add(1, 2))
}
