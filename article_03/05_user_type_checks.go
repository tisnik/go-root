// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 5:
//    Kontrola uživatelsky definovaných typů.

package main

import "fmt"

type Id uint32
type Name string
type Surname string

func main() {
	var i Id
	i = 0
	fmt.Println(i)

	var str = "common string"

	var n Name = str
	var s Surname = str

	n = "Jan"
	s = "Novák"

	fmt.Println(n)
	fmt.Println(s)
}
