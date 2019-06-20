// Seriál "Programovací jazyk Go"
//
// Dvacátá pátá část
//
// Demonstrační příklad číslo 1:
//     	Jednoduchý program, který sečte dvojici celočíselných hodnot

package main

import "fmt"

func main() {
	x := 10
	y := 20

	z := x + y

	fmt.Printf("%d + %d = %d\n", x, y, z)
}
