// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 7:
//    Změna hodnoty již existující proměnné

package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	b := "hello"
	fmt.Println(b)
	c := true
	fmt.Println(c)

	a = 20
	fmt.Println(a)
	b = "world"
	fmt.Println(b)
	c = false
	fmt.Println(c)
}
