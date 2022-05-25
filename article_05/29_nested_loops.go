// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Seznam příkladů z páté části:
//    https://github.com/tisnik/go-root/blob/master/article_05/README.md
//
// Demonstrační příklad číslo 29:
//    Dvojice vnořených smyček for.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_05/29_nested_loops.html

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%3d ", i*j)
		}
		fmt.Println()
	}
}
