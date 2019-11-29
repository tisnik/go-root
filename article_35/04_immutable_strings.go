// Seriál "Programovací jazyk Go"
//
// Třicátá pátá část
//
// Demonstrační příklad číslo 4:
//     Řetězce jsou v Go neměnitelné.

package main

import "fmt"

func main() {
	var s string = "www.root.cz"

	fmt.Println(s)
	s[3] = '*'
	s[8] = '*'
	fmt.Println(s)
}
