// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 11:
//    Zjištění, kdy se vyhodnocují výraz pro volání funkce v defer().

package main

import "fmt"

func function(i int) {
	fmt.Printf("Defer %2d\n", i)
}

func main() {
	x := 0

	fmt.Printf("Current x value = %2d\n", x)
	defer function(x)

	x++

	fmt.Printf("Current x value = %2d\n", x)
	defer function(x)

	x++
	fmt.Printf("Current x value = %2d\n", x)
}
