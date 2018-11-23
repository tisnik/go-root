// Seriál "Programovací jazyk Go"
//
// První část
//
// Demonstrační příklad číslo 8:
//    Špatný typ hodnoty přiřazený do proměnných

package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	b := "hello"
	fmt.Println(b)
	c := true
	fmt.Println(c)

	a = "world"
	fmt.Println(a)
	b = 0
	fmt.Println(b)
	c = nil
	fmt.Println(c)
}
