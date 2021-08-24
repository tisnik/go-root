// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 13:
//    Defer a příkaz return.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/13_defer_on_all_returns.html

package main

import "fmt"

func function(x int) {
	fmt.Printf("Defer %d\n", x)
}

func classify(x int) string {
	defer function(x)
	switch x {
	case 0:
		return "nula"
	case 2, 4, 6, 8:
		return "sudé číslo"
	case 1, 3, 5, 7, 9:
		return "liché číslo"
	default:
		return "?"
	}
}

func main() {
	for x := 0; x <= 10; x++ {
		println(x, classify(x))
	}
}
