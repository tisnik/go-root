// Seriál "Programovací jazyk Go"
//
// Druhá část
//
// Demonstrační příklad číslo 17
//    Datový typ řetězec

package main

import "fmt"

func main() {
	var s1 string = "Hello\nworld!\n"
	var s2 string = `Hello\nworld!\n`

	fmt.Println(s1)
	fmt.Println(s2)
}
