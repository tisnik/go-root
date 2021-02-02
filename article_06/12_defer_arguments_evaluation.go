// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 12:
//    Zjištění, kdy se vyhodnocují výraz pro volání funkce v defer().

package main

import "fmt"

func function(a []int) {
	fmt.Printf("Defer %v\n", a)
}

func main() {
	var x = []int{1, 2, 3}

	fmt.Printf("Current x value = %v\n", x)
	defer function(x)

	x[0] = 0

	fmt.Printf("Current x value = %v\n", x)
	defer function(x)

	x[1] = 0

	fmt.Printf("Current x value = %v\n", x)
	defer function(x)

	x[2] = 0

	fmt.Printf("Current x value = %v\n", x)
}
