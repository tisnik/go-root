// Seriál "Programovací jazyk Go"
//
// Druhá část
//
// Demonstrační příklad číslo 15
//    Datový typ Boolean - kontroly při překladu

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
