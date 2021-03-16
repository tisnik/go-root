// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá pátá část
//     Ladění aplikací v Go s využitím GNU Debuggeru a debuggeru Delve
//     https://www.root.cz/clanky/ladeni-aplikaci-v-go-s-vyuzitim-gnu-debuggeru-a-debuggeru-delve/
//
// Demonstrační příklad číslo 2:
//     Základní operace s poli a s řezy

package main

import "fmt"

func main() {
	var a1 [20]int

	for i := 0; i < len(a1); i++ {
		a1[i] = i + 1
	}

	var slice1 []int = a1[:]
	var slice2 []int = a1[5:15]
	slice3 := slice2[5:]

	fmt.Println(a1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
