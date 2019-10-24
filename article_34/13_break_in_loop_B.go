// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//
// Demonstrační příklad číslo 13:
//    Nekonečná smyčka v jazyce Go.

package main

import "fmt"

func main() {
	for x := 1; true; { /* endless loop */
		fmt.Printf("%d\n", x)
		x++
		if x > 10 {
			break
		}
	}
}
