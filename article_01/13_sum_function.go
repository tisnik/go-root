// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 13:
//    Funkce se dvěma parametry

package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("sum =", sum(1, 2))
}
