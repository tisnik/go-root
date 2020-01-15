// Seriál "Programovací jazyk Go"
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 20:
//    Čtení hodnot z mapy.

package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	fmt.Println(m1)

	m1["nula"] = 0
	m1["jedna"] = 1
	m1["dva"] = 2
	m1["tri"] = 3
	m1["ctyri"] = 4
	m1["pet"] = 5
	m1["sest"] = 6

	fmt.Println(m1)

	value, exist := m1["dva"]

	if exist {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}

	value, exist = m1["sto"]

	if exist {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}
}
