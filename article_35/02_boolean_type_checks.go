// Seriál "Programovací jazyk Go"
//
// Třicátá pátá část
//
// Demonstrační příklad číslo 2:
//     Převod různých hodnot na typ bool.

package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	var c bool = 0
	var d bool = ""

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
