// Seriál "Programovací jazyk Go"
//
// Dvacátá čtvrtá část
//
// Demonstrační příklad číslo 8:
//     	Prázdná cyklická fronta

package main

import (
	"container/ring"
	"fmt"
)

func printRing(r *ring.Ring) {
	length := r.Len()
	for i := 0; i < length; i++ {
		fmt.Println(i, r.Value)
		r = r.Next()
	}
}

func main() {
	r := ring.New(10)
	printRing(r)
}
