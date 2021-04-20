// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá pátá část
//    Programovací jazyk Go pro skalní céčkaře (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/

package main

import "fmt"

func main() {
	var a1 [10]int

	a2 := a1

	fmt.Printf("Pole 1: %v\n", a1)
	fmt.Printf("Pole 2: %v\n", a2)

	for i := 0; i < len(a1); i++ {
		a1[i] = i * 2
	}

	fmt.Printf("Pole 1: %v\n", a1)
	fmt.Printf("Pole 2: %v\n", a2)
}
