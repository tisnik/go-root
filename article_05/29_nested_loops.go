// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 29:
//    Dvojice vnořených smyček for.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%3d ", i*j)
		}
		fmt.Println()
	}
}
