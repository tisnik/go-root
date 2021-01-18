// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 27:
//    Datový typ "funkce s parametry"

package main

import "fmt"

func funkce1(x int, y int) int {
	return x + y
}

func funkce2(x int, y int) int {
	return x * y
}

func main() {
	var a func(int, int) int
	fmt.Println(a)

	a = funkce1
	fmt.Println(a)
	fmt.Println(a(10, 20))

	a = funkce2
	fmt.Println(a)
	fmt.Println(a(10, 20))
}
