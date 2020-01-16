// Seriál "Programovací jazyk Go"
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 26:
//    Datový typ "funkce"

package main

import "fmt"

func funkce1() {
	fmt.Println("funkce1")
}

func funkce2() {
	fmt.Println("funkce2")
}

func main() {
	var a func()
	fmt.Println(a)

	a = funkce1
	fmt.Println(a)
	a()

	a = funkce2
	fmt.Println(a)
	a()
}
