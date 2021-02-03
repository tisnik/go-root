// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 24:
//    Bitové posuny.

package main

import "fmt"

func main() {
	x := 1

	for i := uint(0); i <= 10; i++ {
		fmt.Printf("%d << %2d == %4d\n", x, i, x<<i)
	}

	fmt.Println()

	x = 10000000

	for i := uint(0); i <= 10; i++ {
		fmt.Printf("%d >> %2d == %4d\n", x, i, x>>i)
	}

}
