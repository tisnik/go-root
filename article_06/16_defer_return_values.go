// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 16:
//    Modifikace návratové hodnoty funkce v sekci defer.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/16_defer_return_values.html

package main

import "fmt"

func funkce1() (i int) {
	i = 1
	return
}

func funkce2() (i int) {
	defer func() { i = 2 }()
	return 1
}

func funkce3() (i int) {
	defer func() { i += 2 }()
	return 1
}

func main() {
	fmt.Printf("Návratová hodnota funkce1: %d\n", funkce1())
	fmt.Printf("Návratová hodnota funkce2: %d\n", funkce2())
	fmt.Printf("Návratová hodnota funkce3: %d\n", funkce3())
}
