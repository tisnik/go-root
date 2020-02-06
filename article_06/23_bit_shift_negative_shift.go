// Seriál "Programovací jazyk Go"
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 23:
//    Bitové posuny s negativním druhým operandem.

package main

import "fmt"

func main() {
	x := 1

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d << %2d == %4d\n", x, i, x<<i)
	}

	fmt.Println()

	x = 10000000

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d >> %2d == %4d\n", x, i, x>>i)
	}

}
