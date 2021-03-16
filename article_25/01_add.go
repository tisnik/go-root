// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá pátá část
//     Ladění aplikací v Go s využitím GNU Debuggeru a debuggeru Delve
//     https://www.root.cz/clanky/ladeni-aplikaci-v-go-s-vyuzitim-gnu-debuggeru-a-debuggeru-delve/
//
// Demonstrační příklad číslo 1:
//     Jednoduchý program, který sečte dvojici celočíselných hodnot

package main

import "fmt"

func main() {
	x := 10
	y := 20

	z := x + y

	fmt.Printf("%d + %d = %d\n", x, y, z)
}
